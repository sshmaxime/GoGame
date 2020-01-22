package main

import (
	log "github.com/sirupsen/logrus"
)

func LogUpdateError(err error) {
	log.Error("GameUpdate: " + err.Error())
}
func LogUpdateInfo(info string) {
	log.Info("GameUpdate: " + info)
}
