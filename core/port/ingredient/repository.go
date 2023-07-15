package ingredient

import (
	domain "svc-receipt-luscious/core/domain/ingredient"
)

type (
	Repository interface {
		GetAllListIngredient(ingredientName string) ([]domain.IngredientService, error)
	}
)
