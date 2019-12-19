package main

import (
	"github.com/joho/godotenv"
	"golang_side_project_crud_website/config"
	"golang_side_project_crud_website/routes"
	"log"
	"net/http"
)

func main() {

	// load the project .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// open all route
	routes.Main()

	// open db connection pool and run migration files
	db := config.OpenDatabaseConnectionPool()
	config.MigrateAndPassDatabaseConnectionToModels(db)
	config.InitFirebaseClient()

	// open server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
