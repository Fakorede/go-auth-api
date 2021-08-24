package main

import (
	"database/sql"
	"goauthapi/driver"
	"goauthapi/handlers"
	"goauthapi/router"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

var DB *sql.DB

const PORT = ":5000"

func init() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}

func main() {
	DB = driver.ConnectDB()

	router := router.Router{}
	
	handlers := handlers.Handler{}

	routes := router.Routes(DB, handlers)

	log.Println("Starting server on PORT", PORT)

	log.Fatal(http.ListenAndServe(PORT, routes))
}

