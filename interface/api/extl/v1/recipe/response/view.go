package response

type (
	Response struct {
		ID         string `json:"id"`
		RecipeName string `json:"recipe_name"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
	}
)
