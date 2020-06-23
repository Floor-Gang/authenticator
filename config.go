package main

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
)

type Config struct {
	Token string
	Roles []string
	Port  int
	Guild string
}

func getConfig() Config {
	var err error

	if _, err = os.Stat(configPath); err != nil {
		return genConfig()
	}
	data, err := ioutil.ReadFile(configPath)

	var config = Config{Roles: []string{}, Port: 6969}
	var configMap = make(map[string]interface{})

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, configMap)

	if err != nil {
		panic(err)
	}

	config.Port = configMap["port"].(int)
	config.Token = configMap["token"].(string)
	config.Guild = configMap["guild"].(string)
	for _, roleID := range configMap["roles"].([]interface{}) {
		config.Roles = append(config.Roles, roleID.(string))
	}

	return config
}

func genConfig() Config {
	var err error
	newConfig := Config{
		Roles: []string{"1", "2", "3"},
		Port:  6969,
		Token: "",
		Guild: "718433475828645928",
	}

	output, err := yaml.Marshal(newConfig)
	err = ioutil.WriteFile(configPath, output, 0644)

	if err != nil {
		panic(err)
	}

	return newConfig
}
