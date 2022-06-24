package main

import (
	"github.com/timkelleher/itinerary/pkg/db"
	server "github.com/timkelleher/itinerary/pkg/server"
)

func main() {
	client, err := db.NewDB()
	if err != nil {
		panic("cannot initialize db")
	}
	client.MigrateUp()

	server.Run(uint16(3000))
}
