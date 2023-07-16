package recipe

import (
	"context"
	domain "svc-receipt-luscious/core/domain/recipe"
)

type (
	Repository interface {
		GetAllListRecipe(recipeName string) ([]domain.Recipe, error)
		InsertRecipe(ctx context.Context, inData *domain.Recipe) error
		UpdateRecipe(ctx context.Context, inData *domain.Recipe) error
		DeleteRecipe(ctx context.Context, recipeID string) error
	}
)
