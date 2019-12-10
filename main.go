package main

import (
	"github.com/joho/godotenv"
	"golang_side_project_crud_website/config"
	"golang_side_project_crud_website/routes"
	"log"
	"net/http"
)

// TODO: csrf機制, https://firebase.google.com/docs/auth/admin/manage-cookies 實作
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

	// open server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
