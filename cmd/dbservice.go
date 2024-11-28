package main

import (
	"hlservice-db/cmd/flags"
	"hlservice-db/internal/app/api"
	"log"

	"github.com/BurntSushi/toml"
)

func main() {
	config := api.NewConfig()

	_, err := toml.DecodeFile(flags.ConfigPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := api.Start(config); err != nil {
		log.Fatal(err)
	}
}
