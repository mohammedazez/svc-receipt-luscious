package response

type (
	Response struct {
		ID             string `json:"id"`
		IngredientName string `json:"ingredient_name"`
		RecipeID       string `json:"recipe_id"`
		Quantity       string `json:"quantity"`
		CreatedAt      string `json:"created_at"`
		UpdatedAt      string `json:"updated_at"`
	}
)
