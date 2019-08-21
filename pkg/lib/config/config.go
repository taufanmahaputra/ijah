package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"sync"
)

var cfg Config
var once sync.Once

func GetConfig() Config {
	once.Do(func() {
		var data string
		switch os.Getenv("ENV") {
		case "PRODUCTION":
			data = "./config/ijah.production.toml"
		default:
			data = "./config/ijah.development.toml"
		}

		if _, err := toml.DecodeFile(data, &cfg); err != nil {
			log.Println(err)
		}
	})

	return cfg
}
