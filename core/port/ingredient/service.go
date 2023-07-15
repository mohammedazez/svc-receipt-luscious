package ingredient

import (
	domain "svc-receipt-luscious/core/domain/ingredient"
)

type (
	Service interface {
		List(ingredientName string) ([]domain.IngredientService, error)
	}
)
