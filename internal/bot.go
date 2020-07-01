package internal

import (
	util "github.com/Floor-Gang/utilpkg"
	dg "github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

type Bot struct {
	client     *dg.Session
	config     Config
	configPath string
}

func StartBot(config Config, path string) {
	client, _ := dg.New("Bot " + config.Token)

	bot := Bot{
		client:     client,
		config:     config,
		configPath: path,
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
}

func (bot *Bot) OnMessage(_ *dg.Session, msg *dg.MessageCreate) {
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
			bot.client,
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
