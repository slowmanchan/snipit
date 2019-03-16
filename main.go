package main

import (
	"log"

	"github.com/norman/snipit/db"
)

func main() {
	_, err := db.NewConnection("postgres://:@localhost:5432/normanchan")
	if err != nil {
		log.Fatal(err)
	}
}
