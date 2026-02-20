package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"Docker_postgre_Folderstruct/database"
	"Docker_postgre_Folderstruct/routes"
)

func main() {
	r := mux.NewRouter()

	database.ConnectDB() // now sets global DB

	routes.RegisterRoutes(r)

	log.Println("Server running on 8080 🚀")
	log.Fatal(http.ListenAndServe(":8080", r))
}