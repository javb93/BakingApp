package main

import (
	"BakingApp/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//Test of correct github account
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", api.HandleGet).Methods("GET")
	router.HandleFunc("/recipes", api.HandleGetRecipes).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}
