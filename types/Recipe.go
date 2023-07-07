package types

type Recipe struct {
	ID          string       `json:"id" :"ID"`
	Name        string       `json:"name" :"name"`
	PeopleQty   int          `json:"peopleQty" :"people_Qty"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
}
