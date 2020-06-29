package internel

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Token  string   `yaml:"token"`
	Prefix string   `yaml:"prefix"`
	Roles  []string `yaml:"roles"`
	Port   int      `yaml:"port"`
	Guild  string   `yaml:"guild"`
}

func GetConfig(path string) (config Config) {
	if _, err := os.Stat(path); err != nil {
		genConfig(path)
		log.Fatalln("Please fill out the configuration file.")
	}

	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalln("Failed to read from configuration file. " + err.Error())
	}

	if err = yaml.Unmarshal(data, config); err != nil {
		log.Fatalln("Failed to parse configuration file. " + err.Error())
	}

	return config
}

func genConfig(path string) {
	config := Config{
		Token:  "",
		Prefix: ".admin",
		Roles:  []string{"Role ID 1", "Role ID 2", "Role ID 3"},
		Port:   6969,
		Guild:  "",
	}

	serialized, err := yaml.Marshal(config)

	if err != nil {
		log.Fatal("Failed to serialize config to a file." + err.Error())
	}

	err = ioutil.WriteFile(path, serialized, 0660)

	if err != nil {
		log.Fatalln("Failed to write the default configuration to a file. " + err.Error())
	}
}
