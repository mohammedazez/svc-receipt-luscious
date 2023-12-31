package category

import (
	"context"
	"fmt"
	"log"
	domain "svc-receipt-luscious/core/domain/category"
	port "svc-receipt-luscious/core/port/category"
)

type Service struct {
	repo port.Repository
}

func NewService(repo port.Repository) port.Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) List(categoryName string) ([]domain.Category, error) {

	category, err := s.repo.GetAllListCategory(categoryName)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *Service) Detail(categoryName string) (*domain.Category, error) {

	category, err := s.repo.GetDetailCategory(categoryName)
	if err != nil {
		return nil, err
	}

	fmt.Println(category)

	return category, nil
}

func (s *Service) Insert(form *domain.Category) error {

	ctx := context.Background()

	err := s.repo.InsertCategory(ctx, form)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Update(form *domain.Category) error {
	ctx := context.Background()

	err := s.repo.UpdateCategory(ctx, form)
	if err != nil {
		log.Printf("Failed to update category: %v", err)
		return err
	}

	return nil
}

func (s *Service) Delete(categoryID string) error {
	ctx := context.Background()

	err := s.repo.DeleteCategory(ctx, categoryID)
	if err != nil {
		log.Printf("Failed to delete category: %v", err)
		return err
	}

	return nil
}
