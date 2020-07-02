package internal

import (
	"fmt"
	util "github.com/Floor-Gang/utilpkg"
	dg "github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

type Bot struct {
	Client     *dg.Session
	config     Config
	configPath string
	Features   map[string]*Feature
}

func StartBot(config Config, path string) Bot {
	client, _ := dg.New("Bot " + config.Token)

	bot := Bot{
		Client:     client,
		config:     config,
		configPath: path,
		Features:   make(map[string]*Feature),
	}

	intents := dg.MakeIntent(
		dg.IntentsGuildMessages,
	)
	client.AddHandler(bot.OnReady)
	client.AddHandler(bot.OnMessage)
	client.Identify.Intents = intents

	if err := client.Open(); err != nil {
		util.Report("Failed to connect to Discord. Is the access token correct?", err)
	}

	return bot
}

// Output:
// Current Features:
// **FaQ Manager**: For managing the FaQ Channel
// - .faq add:
//   - .faq add question NEWLINE answer
// ...
func (bot *Bot) getHelp() dg.MessageEmbed {
	embed := dg.MessageEmbed{
		Title:       "Current Features",
		Description: "These are all the current loaded ",
		Color:       0xef2f2f,
		Fields:      []*dg.MessageEmbedField{},
	}

	for _, feature := range bot.Features {
		field := dg.MessageEmbedField{
			Name:   feature.Name,
			Value:  feature.Description,
			Inline: false,
		}
		if len(feature.Commands) > 0 {
			field.Value += "\n"

			for _, command := range feature.Commands {
				field.Value += fmt.Sprintf("**%s %s**: %s\n", feature.CommandPrefix, command.Name,
					command.Description)
				field.Value += fmt.Sprintf(" - `%s ", feature.CommandPrefix)

				length := len(command.Example)
				for i, example := range command.Example {
					if i == 0 {
						field.Value += fmt.Sprintf("%s", example)
					} else {
						field.Value += fmt.Sprintf("<%s>", example)
					}
					if (length - 1) != i {
						field.Value += " "
					}
				}
				field.Value += "`\n"
			}
		}
		embed.Fields = append(embed.Fields, &field)
	}

	return embed
}

func (bot *Bot) OnMessage(_ *dg.Session, msg *dg.MessageCreate) {
	// .help
	if msg.Content == ".help" {
		helpEmbed := bot.getHelp()
		_, _ = bot.Client.ChannelMessageSendEmbed(
			msg.ChannelID,
			&helpEmbed,
		)
		return
	}

	// Ignore bots and messages that don't start with command prefix
	if msg.Author.Bot || !strings.HasPrefix(msg.Content, bot.config.Prefix) {
		return
	}

	// Ignore non guild members
	if len(msg.GuildID) == 0 {
		return
	}

	// args [".admin", "add" || "remove" || "list", ...role names]
	args := strings.Fields(msg.Content)

	if len(args) < 2 {
		return
	}

	isAdmin, _ := hasRole(msg.Member.Roles, bot.config.Roles)

	if !isAdmin {
		_, _ = util.Reply(
			bot.Client,
			msg.Message,
			"You don't have permissions to run this command.",
		)
		return
	}

	switch args[1] {
	case "add":
		if len(args) >= 3 {
			bot.add(msg, args)
		}
	case "remove":
		if len(args) >= 3 {
			bot.remove(msg, args)
		}
	case "list":
		bot.list(msg)
	}
}

func (bot *Bot) OnReady(_ *dg.Session, ready *dg.Ready) {
	log.Printf("Ready as %s\n", ready.User.Username)
}
