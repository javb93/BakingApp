package main

import (
	"BakingApp/api"
	"BakingApp/db"
	"BakingApp/types"
	"log"
)

// Test of correct github account
func main() {
	database, err := db.ConnectDB()
	// AutoMigrate to make sure the Recipe table exists
	err = database.AutoMigrate(&types.Recipe{})
	err = database.AutoMigrate(&types.Ingredient{})

	if err != nil {
		log.Fatalf("Could not migrate the database: %v", err)
	}
	api.StartServer(database)
}
