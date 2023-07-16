package recipe

import (
	domain "svc-receipt-luscious/core/domain/recipe"
)

type (
	Service interface {
		List(recipeName string) ([]domain.Recipe, error)
		Insert(form *domain.Recipe) error
		Update(form *domain.Recipe) error
	}
)
