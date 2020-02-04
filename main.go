package main

import (
	"github.com/GoGame/models"
	"github.com/GoGame/server"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func getConfig(configPath string) (*models.ServerConfig, error) {
	var configAsBytes []byte
	var err error

	config := models.ServerConfig{}
	if configAsBytes, err = ioutil.ReadFile(configPath); err != nil {
		return nil, err

	}
	if err = yaml.Unmarshal(configAsBytes, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func main() {
	var config *models.ServerConfig
	var err error

	// Getting config
	if config, err = getConfig("./config.yaml"); err != nil {
		log.Println(err)
		return
	}

	// Initializing our server
	goGameServer := server.Server{}
	if err = goGameServer.Init(config); err != nil {
		log.Println(err)
		return
	}

	// Start the server
	goGameServer.Start()
}
