package request

type (
	RequestUpdate struct {
		ID         string `json:"-"`
		RecipeName string `json:"recipe_name"`
		CategoryID string `json:"category_id"`
		HowToMake  string `json:"how_to_make"`
	}
)
