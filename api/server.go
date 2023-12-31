package api

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func setupRoutes(router *mux.Router, database *gorm.DB) {
	router.HandleFunc("/", HandleGet).Methods("GET")
	router.HandleFunc("/recipes", HandleGetRecipes(database)).Methods("GET")
	router.HandleFunc("/recipe/{id}", HandleGetIndRecipe(database)).Methods("GET")
	router.HandleFunc("/recipes", HandlePostRecipes(database)).Methods("POST")
	router.HandleFunc("/recipes/{id}", HandlePutRecipes(database)).Methods("PUT")
	router.HandleFunc("/recipes/{id}", HandleDeleteRecipe(database)).Methods("DELETE")

}

func StartServer(database *gorm.DB) {

	router := mux.NewRouter()
	setupRoutes(router, database)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// For testing purposes
func GetRouter(database *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	setupRoutes(router, database)

	return router
}
