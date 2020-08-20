package main

import (
	"github.com/GoGame/config"
	"github.com/GoGame/server"
	"log"
)

func main() {
	var err error

	if err = config.Init("./config.yaml"); err != nil {
		log.Println(err)
		return
	}

	if err = server.Init(); err != nil {
		log.Println(err)
		return
	}

	server.Start()
}
