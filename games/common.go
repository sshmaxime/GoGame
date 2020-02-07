package main

import (
	log "github.com/sirupsen/logrus"
)

func LogUpdateError(err error) error {
	log.Error("GameUpdate: " + err.Error())
	return err
}
func LogUpdateInfo(info string) {
	log.Info("GameUpdate: " + info)
}
