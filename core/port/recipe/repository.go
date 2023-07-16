package recipe

import (
	"context"
	domain "svc-receipt-luscious/core/domain/recipe"
)

type (
	Repository interface {
		GetAllListRecipe(recipeName string) ([]domain.Recipe, error)
		InsertRecipe(ctx context.Context, inData *domain.Recipe) error
		// 	UpdateIngredient(ctx context.Context, inData *domain.IngredientService) error
		// 	DeleteIngredient(ctx context.Context, ingredientID string) error
	}
)
