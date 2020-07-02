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

func (bot *Bot) getHelp() (features string) {
	features = "Current Features:\n"
	for _, feature := range bot.Features {
		features += fmt.Sprintf("**%s**: %s\n", feature.Name, feature.Description)
		for _, command := range feature.Commands {
			features += fmt.Sprintf(" - %s: %s\n", command.Name, command.Description)
			features += " - `"
			for i, example := range command.Example {
				if i == (len(command.Example) - 1) {
					features += fmt.Sprintf("%s, ", example)
				} else {
					features += fmt.Sprintf("%s`\n", example)
				}
			}
		}
	}
	return features
}

func (bot *Bot) OnMessage(_ *dg.Session, msg *dg.MessageCreate) {
	// .help
	if msg.Content == ".help" {
		_, _ = util.Reply(bot.Client, msg.Message, bot.getHelp())
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
