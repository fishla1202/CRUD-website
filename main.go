package main

import (
	"golang_side_project_crud_website/config"
	"golang_side_project_crud_website/routes"
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
