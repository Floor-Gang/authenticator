package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	configPath = "./config.yml"
)

var (
	client      *discordgo.Session
	adminConfig Config
)

type AuthArgs struct{ MemberID string }

type AuthResponse struct {
	Role    string
	IsAdmin bool
}

type AuthServer struct{}

func main() {
	var err error
	adminConfig = getConfig()

	// Start Discord
	client, err = discordgo.New("Bot " + adminConfig.Token)

	if err != nil {
		panic(err)
	}

	err = client.Open()

	if err != nil {
		panic(err)
	}

	// Start the auth server
	authServer := new(AuthServer)

	err = rpc.Register(authServer)

	if err != nil {
		panic(err)
	}

	rpc.HandleHTTP()

	fmt.Printf("Listening on port %d\n", adminConfig.Port)

	err = http.ListenAndServe(":"+strconv.Itoa(adminConfig.Port), nil)

	if err != nil {
		panic(err)
	}

	// keep alive
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-signalChan
}

func (a *AuthServer) Auth(args *AuthArgs, reply *AuthResponse) error {
	member, err := client.GuildMember(adminConfig.Guild, args.MemberID)
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
		hasNeededRole := hasRole(roleID)
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

	log.Println(
		fmt.Sprintf("%s is an admin", args.MemberID),
	)

	return nil
}

func hasRole(role string) bool {
	for _, roleID := range adminConfig.Roles {
		if roleID == role {
			return true
		}
	}
	return false
}
