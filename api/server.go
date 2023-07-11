package api

import (
	"BakingApp/db"
	"BakingApp/types"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func setupRoutes(router *mux.Router, database *gorm.DB) {
	router.HandleFunc("/", HandleGet).Methods("GET")
	router.HandleFunc("/recipes", HandleGetRecipes(database)).Methods("GET")
}

func StartServer() {
	database, err := db.ConnectDB()
	// AutoMigrate to make sure the Recipe table exists
	err = database.AutoMigrate(&types.Recipe{})
	err = database.AutoMigrate(&types.Ingredient{})

	if err != nil {
		log.Fatalf("Could not migrate the database: %v", err)
	}
	router := mux.NewRouter()
	setupRoutes(router, database)

	log.Fatal(http.ListenAndServe(":8080", router))
}
