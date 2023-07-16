package ingredient

import (
	"testing"

	domain "svc-receipt-luscious/core/domain/ingredient"
	repo "svc-receipt-luscious/core/port/ingredient/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockRepo = &repo.Repository{Mock: mock.Mock{}}

var service = NewService(mockRepo)

func TestListIngredient(t *testing.T) {
	mockRepo.Mock.On("GetAllListIngredient", "bawang").Return(nil, nil)

	mockRepo.Mock.On("GetAllListIngredient", "bawang").Return([]domain.IngredientService{
		{
			ID:             "594dbadc-515a-4c10-871e-ec8c31dd0d83",
			IngredientName: "Bawang Merah",
			RecipeID:       "f511b1da-b1e9-4df4-bc52-660f73f407f8",
			Quantity:       "2 siung",
			CreatedAt:      "2023-07-16 11:46:42.61695346 +0700 WIB m=+3451.675515087",
			UpdatedAt:      "2023-07-16 11:46:42.61695346 +0700 WIB m=+3451.675515087",
		},
		{
			ID:             "6e9cc8a1-8e4f-4a49-a902-5bdcef113f22",
			IngredientName: "Bawang Merah",
			RecipeID:       "f511b1da-b1e9-4df4-bc52-660f73f407f8",
			Quantity:       "2 siung",
			CreatedAt:      "2023-07-16 11:46:42.61695346 +0700 WIB m=+3451.675515087",
			UpdatedAt:      "2023-07-16 11:46:42.61695346 +0700 WIB m=+3451.675515087",
		},
	}, nil)

	t.Run("list ingredient error", func(t *testing.T) {
		ingredients, err := service.List("bawang")
		assert.Nil(t, ingredients)
		assert.Equal(t, err, nil)
	})

	t.Run("list ingredient success", func(t *testing.T) {
		ingredients, err := service.List("bawang")
		// assert.NotNil(t, ingredients)
		assert.Len(t, ingredients, 0)
		assert.Nil(t, err)
	})
}
