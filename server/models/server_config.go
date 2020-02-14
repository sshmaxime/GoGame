package models

type ServerConfig struct {
	Port                  string `yaml:"port"`
	MaxSimultaneousPlayer uint   `yaml:"max_simultaneous_player"`

	//
	Games []struct {
		Name    string `yaml:"name"`
		LibPath string `yaml:"lib_path"`
	} `yaml:"games"`
}
