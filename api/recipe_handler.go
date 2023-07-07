package api

import (
	"BakingApp/types"
	"encoding/json"
	"net/http"
)

func HandleGetRecipes(w http.ResponseWriter, r *http.Request) {
	recipe := types.Recipe{
		ID:        "123",
		Name:      "Pancakes",
		PeopleQty: 4,
		Ingredients: []types.Ingredient{
			{Name: "Flour", Quantity: 250},
			{Name: "Milk", Quantity: 500},
			{Name: "Eggs", Quantity: 2},
		},
	}
	jsonData, err := json.Marshal(recipe)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)
}
