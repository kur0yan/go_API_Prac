package handlers

import (
	"fmt"
	"net/http"
)

//This function basically displays a basic text whenever you need to access the '\' endpoint

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You just hit the home endpoint. Why tho? There's nothing here")
	fmt.Fprintf(w, "You just hit the home endpoint. Why tho? There's nothing here")
}
