package api_test

import (
	"BakingApp/api"
	"BakingApp/types"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func dbInit(db *gorm.DB) {
	db.AutoMigrate(&types.Ingredient{}, &types.Recipe{})
}

func TestGetRecipes(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	dbInit(db)
	db.Create(&types.Recipe{
		Name:      "Pancakes",
		PeopleQty: 4,
		Ingredients: []types.Ingredient{
			{Name: "Flour", Quantity: 250},
			{Name: "Milk", Quantity: 500},
			{Name: "Eggs", Quantity: 2},
		},
	})

	req, err := http.NewRequest("GET", "/recipes", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	api.GetRouter(db).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedRecipe := types.Recipe{
		Name:      "Pancakes",
		PeopleQty: 4,
		Ingredients: []types.Ingredient{
			{Name: "Flour", Quantity: 250},
			{Name: "Milk", Quantity: 500},
			{Name: "Eggs", Quantity: 2},
		},
	}

	if err != nil {
		t.Fatal(err)
	}

	var resultRecipes []types.Recipe
	json.Unmarshal(rr.Body.Bytes(), &resultRecipes)

	assert.Equal(t, 1, len(resultRecipes))
	assert.Equal(t, expectedRecipe.Name, resultRecipes[0].Name)
	assert.Equal(t, 3, len(resultRecipes[0].Ingredients))
	assert.Equal(t, expectedRecipe.Ingredients[0].Name, resultRecipes[0].Ingredients[0].Name)
}

func TestHandlePostRecipe(t *testing.T) {
	// Open an SQLite in-memory database for testing
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	// Run database migration to create the recipes table
	dbInit(db)

	// Initialize the mux router with the database connection
	r := api.GetRouter(db)

	// Create a new recipe to send in the request body
	newRecipe := &types.Recipe{Name: "New Recipe"}
	newRecipeJson, _ := json.Marshal(newRecipe)

	// Create a new request
	req, _ := http.NewRequest("POST", "/recipes", bytes.NewBuffer(newRecipeJson))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Serve the request using the mux router
	r.ServeHTTP(rr, req)

	// Check that the response status code is 201 Created
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Check that the new recipe is stored in the database
	var storedRecipe types.Recipe
	db.First(&storedRecipe, "name = ?", newRecipe.Name)
	assert.Equal(t, newRecipe.Name, storedRecipe.Name)
}
