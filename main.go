package main

import (
	"github.com/andrade-felipe-dev/first-api-in-go/database"
	"github.com/andrade-felipe-dev/first-api-in-go/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
