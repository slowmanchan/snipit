package main

import (
	"log"

	"github.com/norman/snipit/db"
	"github.com/norman/snipit/server"
)

var conn *db.Connection

func main() {
	var err error
	conn, err = db.NewConnection("postgres://:@localhost:5432/snipit")
	if err != nil {
		log.Fatal(err)
	}
	if err := server.Start(conn); err != nil {
		log.Fatal(err)
	}
}
