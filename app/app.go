package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/pokemon", handlers.getPokemon).Methods("GET")
	log.Fatal(http.ListenAndServe(":10001", myRouter))
}
