package category

import (
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

	ingredient, err := s.repo.GetAllListCategory(categoryName)
	if err != nil {
		return nil, err
	}

	return ingredient, nil
}
