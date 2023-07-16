package recipe

import (
	domain "svc-receipt-luscious/core/domain/recipe"
)

type (
	Repository interface {
		GetAllListRecipe(recipeName string) ([]domain.Recipe, error)
		// 	InsertIngredient(ctx context.Context, inData *domain.IngredientService) error
		// 	UpdateIngredient(ctx context.Context, inData *domain.IngredientService) error
		// 	DeleteIngredient(ctx context.Context, ingredientID string) error
	}
)
