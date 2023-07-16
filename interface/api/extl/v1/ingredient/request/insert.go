package request

type (
	RequestInsert struct {
		IngredientName string `json:"ingredient_name" validate:"required" required:"Nama bahan makanan harus diisi"`
		RecipeID       string `json:"recipe_id"  validate:"required" required:"recipe id harus diisi"`
		Quantity       string `json:"quantity" validate:"required" required:"kuantitas harus diisi"`
	}
)
