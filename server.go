package main

import (
	"fmt"
	"log"
	"master-golang-auth/routes"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"
)

func init() {

	port := os.Getenv("web_port")

	prefix := os.Getenv("prefix")
	fmt.Println("Server started at " + port + "...")
	r := mux.NewRouter().StrictSlash(true)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	})
	// Routes
	routes.ApiRoutes(prefix, r)
	handler := c.Handler(r)

	//Start Server on the port set in your .env file
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal(err)
	}
}
