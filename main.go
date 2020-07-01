package main

import (
	"github.com/Floor-Gang/authserver/internal"
	util "github.com/Floor-Gang/utilpkg"
)

const (
	configPath = "./config.yml"
)

func main() {
	var config = internal.GetConfig(configPath)

	// Start Discord
	internal.StartBot(config, configPath)

	// Start the auth server
	internal.StartServer(config)

	// Keep auth-server alive
	util.KeepAlive()
}
