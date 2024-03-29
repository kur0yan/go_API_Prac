package app

import (
	"fmt"
	"log"
	"net/http"
	"pokeprac/app/handlers"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", handlers.HandleHome)
	myRouter.HandleFunc("/pokemon", handlers.GetPokeList).Methods("GET")
	myRouter.HandleFunc("/pokemon", handlers.PostPoke).Methods("POST")
	myRouter.HandleFunc("/pokemon/{id}", handlers.PokeNumSearch).Methods("GET")
	fmt.Println("Listening on localhost:10001")
	log.Fatal(http.ListenAndServe(":10001", myRouter))
}
