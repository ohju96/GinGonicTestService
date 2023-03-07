package config

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

type Config struct {
	Server struct {
		LocalServer string
	}

	Db struct {
		Dbms     string
		Db       string
		User     string
		Password string
		Host     string
	}
}

func InitTomlConfig(path string) *Config {

	config := new(Config)

	if file, err := os.Open(path); err != nil {
		return nil
	} else {
		defer file.Close()

		if err := toml.NewDecoder(file).Decode(config); err != nil {
			return nil
		} else {
			return config
		}
	}
}
