package internal

import (
	"fmt"
	"github.com/Floor-Gang/authserver/pkg"
	util "github.com/Floor-Gang/utilpkg"
	"log"
)

// Auth function authentication framework.
func (a *AuthServer) Auth(args *pkg.AuthArgs, reply *pkg.AuthResponse) error {
	member, err := a.client.GuildMember(a.config.Guild, args.MemberID)

	log.Println(
		fmt.Sprintf("Looking up %s", args.MemberID),
	)

	if err != nil {
		util.Report("Failed to lookup "+args.MemberID, err)
		return err
	}

	isAdmin, role := hasRole(member.Roles, a.config.Roles)

	*reply = pkg.AuthResponse{
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
