package request

type (
	CategoryInsert struct {
		CategoryName string `json:"category_name" validate:"required" required:"kategory makanan harus diisi"`
	}
)
