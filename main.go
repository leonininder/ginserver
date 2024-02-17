package main

import (
	"ginserver/api"
	"ginserver/bin"
	"log"
)

func main() {
	config, err := bin.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	//add server, ex. postgresql, sql ...
	//server, err := api.NewServer(config ...any)
	server, err := api.NewServer(config, false)

	if err != nil {
		log.Fatal("cannot Create server:", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
