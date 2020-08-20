package server

import "github.com/GoGame/database"

func initDatabases() error {
	if err := database.InitUserDatabase(); err != nil {
		return err
	}
	if err := database.InitRoomDatabase(); err != nil {
		return err
	}
	if err := database.InitGameDatabase(); err != nil {
		return err
	}
	return nil
}
