package response

type (
	Response struct {
		ID         string `json:"id"`
		RecipeName string `json:"recipe_name"`
		CategoryID string `json:"category_id"`
		HowToMake  string `json:"how_to_make"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
	}
)
