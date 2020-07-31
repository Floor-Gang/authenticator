package internal

import (
	"github.com/Floor-Gang/authserver/pkg"
	dg "github.com/bwmarrin/discordgo"
	"net/http"
	"net/rpc"
	"strconv"
)

// AuthServer structure.
type AuthServer struct {
	config   Config
	client   *dg.Session             // Discord bot
	Features map[string]*pkg.Feature // Registered features
}

func StartServer(config Config, bot Bot) {
	authServer := new(AuthServer)
	authServer.config = config
	authServer.client = bot.Client
	authServer.Features = bot.Features

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
