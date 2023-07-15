package ingredient

import "time"

type (
	IngredientService struct {
		ID             string
		IngredientName string
		RecipeID       string
		Quantity       string
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
)
