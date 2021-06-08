package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"../models"
	"go.mongodb.org/mongo-driver/bson"
)

func getPokemon(w http.ResponseWriter, r *http.Request) {
	var pokeList []models.Pokemon

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, _ := config.connectDB("mongodb://localhost:27017")
	pokemonCollection = client.Database("gotest").Collection("pokemon")

	cursor, err := pokemonCollection.Find(ctx, bson.M{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var eachPoke models.Pokemon
		cursor.Decode(&eachPoke)
		pokeList = append(pokeList, eachPoke)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(w).Encode(pokeList)
}
