package main

import (
	"log"
	"net/http"
)

type application struct{}

func main() {
	app := application{}

	mux := app.routes()

	log.Println("Starting Server on Port 4444...")

	err := http.ListenAndServe(":4444", mux)
	if err != nil {
		log.Fatal(err)
	}
}
