package internel

import (
	"fmt"
	dg "github.com/bwmarrin/discordgo"
	"log"
)

type AuthArgs struct{ MemberID string }

type AuthResponse struct {
	Role    string
	IsAdmin bool
}

type AuthServer struct {
	config *Config
	client *dg.Session
}

func (a *AuthServer) Auth(args *AuthArgs, reply *AuthResponse) error {
	member, err := a.client.GuildMember(a.config.Guild, args.MemberID)
	isAdmin := false
	role := ""

	log.Println(
		fmt.Sprintf("Looking up %s", args.MemberID),
	)

	if err != nil {
		log.Println(err)
		return err
	}

	for _, roleID := range member.Roles {
		hasNeededRole := a.hasRole(roleID)
		if hasNeededRole {
			isAdmin = true
			role = roleID
			break
		}
	}

	*reply = AuthResponse{
		IsAdmin: isAdmin,
		Role:    role,
	}

	if isAdmin {
		log.Println(
			fmt.Sprintf("%s is an admin", args.MemberID),
		)
	}
	return nil
}

func (a *AuthServer) hasRole(role string) bool {
	for _, roleID := range a.config.Roles {
		if roleID == role {
			return true
		}
	}
	return false
}
