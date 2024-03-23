package main

import (
	"log"

	"github.com/phillipmugisa/MIASoft/api"
	"github.com/phillipmugisa/MIASoft/storage"
)

func main() {
	port, err := GetPort()
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	a := api.NewAppServer(port, db)

	a.StartServer()
}
