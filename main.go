package main

import (
	"CRUD/config"
	"CRUD/routes"
	"log"
	"net/http"
)

func main() {

	// open all route
	routes.Main()

	// open db connection pool and run migration files
	db := config.OpenDatabaseConnectionPool()
	config.MigrateAndPassDatabaseConnectionToModels(db)

	// open server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
