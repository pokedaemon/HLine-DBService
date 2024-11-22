package main

import (
	"hlservice-db/cmd/flags"
	apiserver "hlservice-db/internal/app/api"
	"log"

	"github.com/BurntSushi/toml"
)

func main() {
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(flags.ConfigPath, config)
	if err != nil {
		log.Fatal(err)
	}

	server := apiserver.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
