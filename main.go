package main

import (
	"log"
	"net/http"
	"os"

	"github.com/armr-dev/cypher-api-go/routers"
)

func main() {
	var port string
	router := Routers.Route()

	_, envExists := os.LookupEnv("PORT")

	if envExists {
		port = os.Getenv("PORT")
	} else {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}
