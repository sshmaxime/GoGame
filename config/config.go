package config

import (
	"github.com/GoGame/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var config *ServerConfig

type ServerConfig struct {
	Port                  string `yaml:"port"`
	MaxSimultaneousPlayer uint   `yaml:"max_simultaneous_player"`

	Games []types.GameDefinition `yaml:"games"`
}

func Init(configPath string) (err error) {
	var configAsBytes []byte

	if configAsBytes, err = ioutil.ReadFile(configPath); err != nil {
		return err

	}
	if err = yaml.Unmarshal(configAsBytes, &config); err != nil {
		return err
	}
	return nil
}

func GetGames() []types.GameDefinition {
	return config.Games
}

func GetPort() string {
	return config.Port
}
