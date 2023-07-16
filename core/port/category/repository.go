package category

import (
	domain "svc-receipt-luscious/core/domain/category"
)

type (
	Repository interface {
		GetAllListCategory(categoryName string) ([]domain.Category, error)
		// InsertIngredient(ctx context.Context, inData *domain.IngredientService) error
		// UpdateIngredient(ctx context.Context, inData *domain.IngredientService) error
		// DeleteIngredient(ctx context.Context, ingredientID string) error
	}
)
