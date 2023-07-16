package response

type (
	Response struct {
		ID           string `json:"id"`
		CategoryName string `json:"category_name"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
	}
)
