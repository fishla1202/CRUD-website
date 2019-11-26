package main

import (
	"CRUD/config"
	"CRUD/models/migrations"
	"CRUD/router"
	"log"
	"net/http"
)

func main() {
	router.Main()
	db := config.DatabaseInit()
	migrations.InitPostTable(db)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
