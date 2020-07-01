package internal

import (
	"fmt"
	util "github.com/Floor-Gang/utilpkg"
	"log"
	"net/http"
	"net/rpc"
	"strconv"
)

func StartServer(config Config) {
	authServer := new(AuthServer)

	err := rpc.Register(authServer)

	if err != nil {
		panic(err)
	}

	rpc.HandleHTTP()

	err = http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)

	if err != nil {
		panic(err)
	}
}

// Auth function authentication framework.
func (a *AuthServer) Auth(args *AuthArgs, reply *AuthResponse) error {
	member, err := a.client.GuildMember(a.config.Guild, args.MemberID)

	log.Println(
		fmt.Sprintf("Looking up %s", args.MemberID),
	)

	if err != nil {
		util.Report("Failed to lookup "+args.MemberID, err)
		return err
	}

	isAdmin, role := hasRole(member.Roles, a.config.Roles)

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

func hasRole(has []string, required []string) (bool, string) {
	for _, role := range has {
		for _, reqRole := range required {
			if reqRole == role {
				return true, ""
			}
		}
	}
	return false, ""
}
