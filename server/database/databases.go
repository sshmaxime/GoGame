package database

func Init() error {
	if err := initUserDatabase(); err != nil {
		return err
	}
	if err := initRoomDatabase(); err != nil {
		return err
	}
	if err := initGameDatabase(); err != nil {
		return err
	}
	return nil
}
