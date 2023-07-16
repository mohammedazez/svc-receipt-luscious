package category

import (
	"testing"

	domain "svc-receipt-luscious/core/domain/category"
	repo "svc-receipt-luscious/core/port/category/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockRepo = &repo.Repository{Mock: mock.Mock{}}

var service = NewService(mockRepo)

func TestListCategory(t *testing.T) {
	mockRepo.Mock.On("GetAllListCategory", "utama").Return(nil, nil)

	mockRepo.Mock.On("GetAllListCategory", "utama").Return([]domain.Category{
		{
			ID:           "594dbadc-515a-4c10-871e-ec8c31dd0d83",
			CategoryName: "Menu Utama",
			CreatedAt:    "2023-07-16 11:46:42.61695346 +0700 WIB m=+3451.675515087",
			UpdatedAt:    "2023-07-16 11:46:42.61695346 +0700 WIB m=+3451.675515087",
		},
		{
			ID:           "6e9cc8a1-8e4f-4a49-a902-5bdcef113f22",
			CategoryName: "Masakan Sayuran",
			CreatedAt:    "2023-07-16 11:46:42.61695346 +0700 WIB m=+3451.675515087",
			UpdatedAt:    "2023-07-16 11:46:42.61695346 +0700 WIB m=+3451.675515087",
		},
	}, nil)

	t.Run("list category error", func(t *testing.T) {
		categories, err := service.List("utama")
		assert.Nil(t, categories)
		assert.Equal(t, err, nil)
	})

	t.Run("list category success", func(t *testing.T) {
		categories, err := service.List("utama")
		// assert.NotNil(t, categories)
		assert.Len(t, categories, 0)
		assert.Nil(t, err)
	})
}
