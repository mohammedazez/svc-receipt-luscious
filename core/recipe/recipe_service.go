package recipe

import (
	"context"
	"log"
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

func (s *Service) Insert(form *domain.Recipe) error {

	ctx := context.Background()

	err := s.repo.InsertRecipe(ctx, form)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Update(form *domain.Recipe) error {
	ctx := context.Background()

	err := s.repo.UpdateRecipe(ctx, form)
	if err != nil {
		log.Printf("Failed to update recipe: %v", err)
		return err
	}

	return nil
}
