package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"pokeprac/app/models"
	"pokeprac/config"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func PokeNumSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var pokeList []models.Pokemon

	parameters := mux.Vars(r)
	id := parameters["id"]

	pokenum, _ := strconv.Atoi(id)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client := config.ConnectDB("mongodb://localhost:27017")
	pokemonCollection := client.Database("gotest").Collection("pokemon")

	cursor, err := pokemonCollection.Find(ctx, bson.M{"pokedex_number": pokenum})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	defer cursor.Close(ctx)

	if cursor.All(context.TODO(), &pokeList); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		fmt.Println("Your dumbass made an error")
		return
	}

	fmt.Println(pokeList)

	json.NewEncoder(w).Encode(pokeList)

}
