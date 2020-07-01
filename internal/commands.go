package internal

import (
	"fmt"
	util "github.com/Floor-Gang/utilpkg"
	dg "github.com/bwmarrin/discordgo"
)

func (bot *Bot) add(msg *dg.MessageCreate, args []string) {
	roleIDs := args[2:]
	var response = "Added Roles:\n"
	guildRoles, _ := bot.client.GuildRoles(msg.GuildID)

	for _, roleID := range roleIDs {
		role := getRole(roleID, guildRoles)
		if role == nil {
			_, _ = util.Reply(bot.client, msg.Message, "Couldn't find "+roleID)
			continue
		}
		response += fmt.Sprintf(" - %s\n", role.Name)
		bot.config.Roles = append(bot.config.Roles, roleID)
	}

	err := save(bot.config, bot.configPath)

	if err != nil {
		_, _ = util.Reply(bot.client, msg.Message, "Failed to remove roles.")
		util.Report("Failed to remove roles", err)
	} else {
		_, _ = util.Reply(bot.client, msg.Message, response)
	}
}

// args = [".admin", "remove" .. role ID's].
func (bot *Bot) remove(msg *dg.MessageCreate, args []string) {
	// ["Discord mod", "developer", "
	roleIDs := args[2:]
	var newConfigRoles []string
	var response = "Removed Roles:\n"
	guildRoles, _ := bot.client.GuildRoles(msg.GuildID)

	for _, storedRole := range bot.config.Roles {
		for _, roleID := range roleIDs {
			role := getRole(roleID, guildRoles)
			if role == nil {
				_, _ = util.Reply(bot.client, msg.Message, "Couldn't find "+roleID)
				continue
			}
			if storedRole == roleID {
				response += fmt.Sprintf(" - %s\n", role.Name)
			} else {
				newConfigRoles = append(newConfigRoles, roleID)
			}
		}
	}

	bot.config.Roles = newConfigRoles

	err := save(bot.config, bot.configPath)

	if err != nil {
		_, _ = util.Reply(bot.client, msg.Message, "Failed to remove roles.")
		util.Report("Failed to remove roles", err)
	} else {
		_, _ = util.Reply(bot.client, msg.Message, response)
	}
}

func getRole(roleID string, roles []*dg.Role) *dg.Role {
	for _, role := range roles {
		if role.ID == roleID {
			return role
		}
	}
	return nil
}

func (bot *Bot) list(msg *dg.MessageCreate) {
	list := fmt.Sprintf("<@%s> Admin Roles:\n", msg.Author.ID)
	roles, err := bot.client.GuildRoles(bot.config.Guild)

	if err != nil {
		util.Report("Failed to fetch roles for "+bot.config.Guild, err)
		_, _ = util.Reply(bot.client, msg.Message, "Failed. See logs.")
		return
	}

	for _, role := range roles {
		if has(role.ID, bot.config.Roles) {
			list += fmt.Sprintf(" - %s\n", role.Name)
		}
	}

	_, _ = bot.client.ChannelMessageSend(msg.ChannelID, list)
}

func has(x string, y []string) bool {
	for _, z := range y {
		if x == z {
			return true
		}
	}
	return false
}
