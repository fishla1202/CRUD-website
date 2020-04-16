package main

import (
	_ "github.com/joho/godotenv/autoload"
	_ "golang_side_project_crud_website/config"
	_ "golang_side_project_crud_website/models"
	"golang_side_project_crud_website/routes"
	"log"
	"net/http"
)

func main() {
	// open all route
	routes.Main()
	// open server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
