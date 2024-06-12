package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ragazoni/consulta-temperatura/api/handlers"
)

func main() {
	http.HandleFunc("/weather", handlers.GetTemperatureHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
