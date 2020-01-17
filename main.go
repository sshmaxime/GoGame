package main

import (
	"github.com/GoGame/server"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	NetworkConfig server.NetworkConfig `yaml:"network_config"`
	ServerConfig  server.ServerConfig  `yaml:"server_config"`
}

func getConfig() (*Config, error) {
	config := Config{}

	configAsBytes, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(configAsBytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func main() {
	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	goGameServer := server.Server{}
	err = goGameServer.Init(&config.NetworkConfig, &config.ServerConfig)
	if err != nil {
		log.Fatal(err)
	}

	goGameServer.Start()
}
