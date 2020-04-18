package configuration

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

func LoadConfiguration() (Config, error) {
	var conf Config
	if _, err := toml.DecodeFile("./configuration/config.toml", &conf); err != nil {
		fmt.Println(err)
	}
	return conf, nil
}

type Config struct {
	Database database
	Api      api
}

type database struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	DbDriver string
}

type api struct {
	Host string
	Port string
}
