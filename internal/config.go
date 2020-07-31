package internal

import (
	util "github.com/Floor-Gang/utilpkg/config"
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

var location *string

// GetConfig retrieves config as Config from path.
func GetConfig(path string) (config Config) {
	location = &path
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

func (config *Config) save() error {
	return util.Save(*location, config)
}
