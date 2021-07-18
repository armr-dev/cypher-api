package Routers

import (
	"github.com/gorilla/mux"
	"github.com/armr-dev/cypher-api-go/controllers"
)

func Route() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Controllers.HomePage)
	router.HandleFunc("/cypher", Controllers.CypherText).Methods("POST", "OPTIONS")
	router.HandleFunc("/decipher", Controllers.DecipherText).Methods("POST", "OPTIONS")

	return router
}