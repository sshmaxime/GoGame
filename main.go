package main

import (
	"fmt"
	"github.com/GoGame/network"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	NetworkConfig network.NetworkConfig `yaml:"network_config"`
	ServerConfig  network.ServerConfig  `yaml:"server_config"`
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

	fmt.Println(config)

	server := network.Server{}
	server.Init(&config.NetworkConfig, &config.ServerConfig)

	server.Start()
}
