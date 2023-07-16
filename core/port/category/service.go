package category

import (
	domain "svc-receipt-luscious/core/domain/category"
)

type (
	Service interface {
		List(categoryName string) ([]domain.Category, error)
		// Detail(categoryName string) (*domain.Category, error)
		Insert(form *domain.Category) error
		Update(form *domain.Category) error
		Delete(categoryID string) error
	}
)
