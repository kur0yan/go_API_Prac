package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"pokeprac/app/models"
	"pokeprac/config"
	"time"
)

func PostPoke(w http.ResponseWriter, r *http.Request) {
	var poke models.Pokemon

	client := config.ConnectDB("mongodb://localhost:27017")

	pokemonCollection := client.Database("gotest").Collection("pokemon")

	r.Header.Add("content-type", "application/json")

	_ = json.NewDecoder(r.Body).Decode(&poke)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, _ := pokemonCollection.InsertOne(ctx, poke)

	print(res)
	json.NewEncoder(w).Encode(res)

}
