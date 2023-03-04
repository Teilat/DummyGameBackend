package main

import (
	"DummyGameBackend/internal/config"
	"DummyGameBackend/internal/db"
	"DummyGameBackend/webapi"
)

func main() {
	config.GetConf()

	db := db.NewDbProvider()
	database, err := db.StartDatabase()
	if err != nil {
		return
	}

	w := webapi.NewWebapi(database)
	w.Start()
}
