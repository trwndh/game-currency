package config

import (
	"log"

	configReader "github.com/trwndh/game-currency/pkg/config"
	"github.com/trwndh/game-currency/pkg/env"
)

type MainConfig struct {
	Server struct {
		Port    string
		BaseURL string
	}

	Database struct {
		MasterDSN     string
		SlaveDSN      string
		RetryInterval int
		MaxIdleConn   int
		MaxConn       int
	}

	Secret struct {
		SecretKey string
	}
}

// ReadConfig : function to read secret file
func ReadConfig(cfg interface{}, module string) interface{} {
	var err error
	if env.IsDevelopment() {
		err = configReader.ReadModuleConfigWithErr(cfg, "files/etc/", module)
	} else {
		err = configReader.ReadModuleConfigWithErr(cfg, "/opt/files/etc", module)
	}
	if err != nil {
		log.Fatalln("Failed To Read Config: ", err)
	}
	return cfg
}

func LoadMainConfig() (cfg *MainConfig) {
	cfg = &MainConfig{}
	ReadConfig(cfg, "main")
	return cfg
}
