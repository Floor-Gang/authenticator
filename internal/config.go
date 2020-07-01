package internal

import (
	util "github.com/Floor-Gang/utilpkg"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"strings"
)

// Config structure.
type Config struct {
	Token  string   `yaml:"token"`
	Prefix string   `yaml:"prefix"`
	Roles  []string `yaml:"roles"`
	Port   int      `yaml:"port"`
	Guild  string   `yaml:"guild"`
}

// GetConfig retrieves config as Config from path.
func GetConfig(path string) (config Config) {
	config = Config{
		Token:  "",
		Prefix: ".admin",
		Roles:  []string{"1", "2", "3"},
		Port:   6969,
		Guild:  "",
	}

	err := util.GetConfig(path, &config)

	if err != nil {
		if strings.Contains(err.Error(), "default") {
			log.Fatalln("A default configuration has been created")
		} else {
			panic(err)
		}
	}

	return config
}

func save(config Config, path string) error {
	serialized, _ := yaml.Marshal(&config)
	err := ioutil.WriteFile(path, serialized, 0660)

	if err != nil {
		util.Report("Failed to save configuration", err)
	}
	return err
}
