package ingredient

import (
	"context"
	"log"
	domain "svc-receipt-luscious/core/domain/ingredient"
	port "svc-receipt-luscious/core/port/ingredient"
)

type Service struct {
	repo port.Repository
}

func NewService(repo port.Repository) port.Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) List(ingredientName string) ([]domain.IngredientService, error) {

	ingredient, err := s.repo.GetAllListIngredient(ingredientName)
	if err != nil {
		return nil, err
	}

	return ingredient, nil
}

func (s *Service) Insert(form *domain.IngredientService) error {

	ctx := context.Background()

	err := s.repo.InsertIngredient(ctx, form)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Update(form *domain.IngredientService) error {
	ctx := context.Background()

	err := s.repo.UpdateIngredient(ctx, form)
	if err != nil {
		log.Printf("Failed to update ingredient: %v", err)
		return err
	}

	return nil
}

func (s *Service) Delete(ingredientID string) error {
	ctx := context.Background()

	err := s.repo.DeleteIngredient(ctx, ingredientID)
	if err != nil {
		log.Printf("Failed to delete ingredient: %v", err)
		return err
	}

	return nil
}
