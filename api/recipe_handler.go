package api

import (
	"BakingApp/types"
	"encoding/json"
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
