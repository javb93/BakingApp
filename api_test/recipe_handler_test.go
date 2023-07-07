package api_test

import (
	"BakingApp/api"
	"BakingApp/types"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func server() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/recipes", api.HandleGetRecipes).Methods("GET")
	return router
}

func TestGetRecipes(t *testing.T) {
	req, err := http.NewRequest("GET", "/recipes", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	server().ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expectedRecipe := types.Recipe{
		ID:        "123",
		Name:      "Pancakes",
		PeopleQty: 4,
		Ingredients: []types.Ingredient{
			{Name: "Flour", Quantity: 250},
			{Name: "Milk", Quantity: 500},
			{Name: "Eggs", Quantity: 2},
		},
	}
	// Marshal the expected Recipe to JSON
	expectedJSON, err := json.Marshal(expectedRecipe)
	if err != nil {
		t.Fatal(err)
	}

	// Compare the expected JSON with the actual response body
	if !reflect.DeepEqual(rr.Body.String(), string(expectedJSON)) {
		t.Errorf("expected JSON %s but got %s", expectedJSON, rr.Body.String())
	}
}
