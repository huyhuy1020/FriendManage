package main

import (
	"Assignment/database"
	"Assignment/handler"
	"log"
	"net/http"
)

func main() {
	var database, err = database.Initialize()
	if err != nil {
		log.Fatalf("could not set up database %v", err)
	}
	defer database.Conn.Close()
	handlers := handler.NewHandler(database)
	log.Println("server start on 8080: http://localhost:8080")
	http.ListenAndServe(":8080", handlers)
}
