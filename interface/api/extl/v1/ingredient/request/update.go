package request

type (
	RequestUpdate struct {
		ID             string `json:"-"`
		IngredientName string `json:"ingredient_name"`
		RecipeID       string `json:"recipe_id"`
		Quantity       string `json:"quantity"`
	}
)
