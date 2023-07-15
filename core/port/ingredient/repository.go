package ingredient

import (
	"context"
	domain "svc-receipt-luscious/core/domain/ingredient"
)

type (
	Repository interface {
		GetAllListIngredient(ingredientName string) ([]domain.IngredientService, error)
		InsertIngredient(ctx context.Context, inData *domain.IngredientService) error
		UpdateIngredient(ctx context.Context, inData *domain.IngredientService) error
	}
)
