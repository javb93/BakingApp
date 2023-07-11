package types

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name        string       `json:"name" :"name"`
	PeopleQty   int          `json:"peopleQty" :"peopleQty"`
	Ingredients []Ingredient `gorm:"many2many:recipe_ingredients" json:"ingredients"`
}

type Ingredient struct {
	gorm.Model
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
}
