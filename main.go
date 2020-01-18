package main

import (
	"github.com/GoGame/models"
	"github.com/GoGame/server"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func getConfig() (*models.ServerConfig, error) {
	config := models.ServerConfig{}

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

	err = goGameServer.Init(config)
	if err != nil {
		log.Fatal(err)
	}

	goGameServer.Start()
}
