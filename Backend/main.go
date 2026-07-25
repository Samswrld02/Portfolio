package main

import (
	"log"

	db "github.com/Samswrld02/Portfolio/db"
	routing "github.com/Samswrld02/Portfolio/routing"
)

func main() {
	//setting router with port for server
	router := routing.NewRouter(db.NewDb())

	//starting server
	if err := router.Start(); err != nil {
		router.Server.Logger.Error("failed to start server", "error", err)
	}

	log.Println("testing")

}
