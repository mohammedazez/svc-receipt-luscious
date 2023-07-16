package ingredient

import (
	domain "svc-receipt-luscious/core/domain/ingredient"
)

type (
	Service interface {
		List(ingredientName string) ([]domain.IngredientService, error)
		Insert(form *domain.IngredientService) error
		Update(form *domain.IngredientService) error
		Delete(ingredientID string) error
	}
)
