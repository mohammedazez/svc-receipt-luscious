package response

import "time"

type (
	Response struct {
		ID             string    `json:"id"`
		IngredientName string    `json:"ingredient_name"`
		RecipeID       string    `json:"recipe_id"`
		Quantity       string    `json:"quantity"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}
)
