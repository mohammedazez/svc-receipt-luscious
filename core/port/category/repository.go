package category

import (
	"context"
	domain "svc-receipt-luscious/core/domain/category"
)

type (
	Repository interface {
		GetAllListCategory(categoryName string) ([]domain.Category, error)
		// GetDetailCategory(categoryID string) (domain.Category, error)
		InsertCategory(ctx context.Context, inData *domain.Category) error
		UpdateCategory(ctx context.Context, inData *domain.Category) error
		// DeleteIngredient(ctx context.Context, ingredientID string) error
	}
)
