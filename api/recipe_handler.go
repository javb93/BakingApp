package api

import (
	"BakingApp/types"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func HandleGetRecipes(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var recipes []types.Recipe
		result := db.Model(&types.Recipe{}).Preload("Ingredients").Find(&recipes)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(recipes)
	}
}
func HandleGetIndRecipe(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var recipe types.Recipe
		if err := db.First(&recipe, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, "Recipe not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		// Respond with the recipe
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(recipe)
	}
}

func HandlePostRecipes(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var recipe types.Recipe
		if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := db.Create(&recipe).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(recipe)
	}

}
func HandlePutRecipes(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var recipe types.Recipe
		if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := db.Model(&types.Recipe{}).Where("id = ?", id).Updates(recipe).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond with the updated recipe
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(recipe)
	}

}
func HandleDeleteRecipe(db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the recipe ID from the URL path
		vars := mux.Vars(r)
		id := vars["id"]

		// Delete the recipe from the database
		if err := db.Delete(&types.Recipe{}, id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond with a 204 No Content status code
		w.WriteHeader(http.StatusNoContent)
	}
}
