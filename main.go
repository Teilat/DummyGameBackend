package main

import (
	"DummyGameBackend/internal/config"
	"DummyGameBackend/internal/db"
	"DummyGameBackend/webapi"
	"time"
)

func main() {
	time.Sleep(10 * time.Second)
	config.GetConf()

	db := db.NewDbProvider(3)
	database, err := db.StartDatabase()
	if err != nil {
		panic(err)
	}

	w := webapi.NewWebapi(database)
	w.Start()
}
