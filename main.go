package main

import (
	. "github.com/Floor-Gang/authserver/internel"
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

func main() {
	var config = GetConfig(configPath)

	// Start Discord
	startBot(config)

	// Start the auth server
	startRPC(config)

	// Keep authserver alive
	keepAlive()
}

func startRPC(config Config) {
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

func startBot(config Config) {
	client, _ := discordgo.New("Bot " + config.Token)
	err := client.Open()

	if err != nil {
		log.Fatalln("Failed to connect to Discord. Is the access token correct?")
	}
}

func keepAlive() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-signalChan
}
