package types

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Conf struct {
	Casdoor struct {
		Endpoint         string
		GrantType        string
		ClientId         string
		ClientSecret     string
		Certificate      string
		OrganizationName string
		ApplicationName  string
		RedirectUri      string
	}
}

var Cfg = Conf{}

func Init() {
	if _, err := toml.DecodeFile("./conf/conf.toml", &Cfg); err != nil {
		panic(err)
	}
	content, err := os.ReadFile(Cfg.Casdoor.Certificate)
	if err != nil {
		panic(err)
	}
	Cfg.Casdoor.Certificate = string(content)
}
