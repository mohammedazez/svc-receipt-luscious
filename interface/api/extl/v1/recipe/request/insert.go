package request

type (
	RequestInsert struct {
		RecipeName string `json:"recipe_name" validate:"required" required:"judul resep makanan harus diisi"`
		CategoryID string `json:"category_id" validate:"required" required:"kategori makanan harus diisi"`
		HowToMake  string `json:"how_to_make" validate:"required" required:"cara pembuatan harus diisi"`
	}
)
