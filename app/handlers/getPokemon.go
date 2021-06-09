package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"pokeprac/app/models"
	"pokeprac/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPokeList(w http.ResponseWriter, r *http.Request) {
	var pokeList []models.Pokemon

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := config.ConnectDB("mongodb://localhost:27017")
	pokemonCollection := client.Database("gotest").Collection("pokemon")

	cursor, err := pokemonCollection.Find(ctx, bson.D{primitive.E{Key: "pokedex_number", Value: 150}})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	defer cursor.Close(ctx)

	if err := cursor.All(context.TODO(), &pokeList); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		fmt.Println("Your dumbass made an error")
		return
	}

	fmt.Println(len(pokeList))
	fmt.Println(pokeList)

	json.NewEncoder(w).Encode(pokeList)
}
