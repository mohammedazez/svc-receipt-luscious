package request

type (
	RequestUpdate struct {
		ID           string `json:"-"`
		CategoryName string `json:"category_name"`
	}
)
