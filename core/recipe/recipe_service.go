package recipe

import (
	domain "svc-receipt-luscious/core/domain/recipe"
	port "svc-receipt-luscious/core/port/recipe"
)

type Service struct {
	repo port.Repository
}

func NewService(repo port.Repository) port.Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) List(recipeName string) ([]domain.Recipe, error) {

	recipe, err := s.repo.GetAllListRecipe(recipeName)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}
