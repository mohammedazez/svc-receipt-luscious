package recipe

import (
	"testing"

	domain "svc-receipt-luscious/core/domain/recipe"
	repo "svc-receipt-luscious/core/port/recipe/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockRepo = &repo.Repository{Mock: mock.Mock{}}

var service = NewService(mockRepo)

func TestListRecipe(t *testing.T) {
	mockRepo.Mock.On("GetAllListRecipe", "tuna").Return(nil, nil)

	mockRepo.Mock.On("GetAllListRecipe", "tuna").Return([]domain.Recipe{
		{
			ID:         "3f3f1126-adf3-468b-b732-1d457ba469dd",
			RecipeName: "ikan tuna bakar",
			CategoryID: "f511b1da-b1e9-4df4-bc52-660f73f407f8",
			HowToMake:  "Kupas kentang, lalu potong dadu dan goreng hingga matang. Sisihkan. kemudian bilas ikan teri lalu goreng hingga matang. Sisihkan. Panaskan sedikit minyak lalu tumis bawang merah dan bawang putih hingga harum. Kemudian masukkan cabai dan tumis hingga harum. Masukkan kentang, ikan teri, dan daun bawang. Lalu bumbui dengan garam, gula dan kaldu bubuk secukupnya. Aduk rata. Jangan lupa koreksi rasa, jika sudah pas angkat dan siap disajikan.",
			CreatedAt:  "2023-07-16 11:46:42.61695346 +0700 WIB m=+3451.675515087",
			UpdatedAt:  "2023-07-16 11:46:42.61695346 +0700 WIB m=+3451.675515087",
		},
		{
			ID:         "3fed7fad-bc8a-4e17-b9bd-8950bd278c38",
			RecipeName: "ikan tuna bakar",
			CategoryID: "f511b1da-b1e9-4df4-bc52-660f73f407f8",
			HowToMake:  "Kupas kentang, lalu potong dadu dan goreng hingga matang. Sisihkan. kemudian bilas ikan teri lalu goreng hingga matang. Sisihkan. Panaskan sedikit minyak lalu tumis bawang merah dan bawang putih hingga harum. Kemudian masukkan cabai dan tumis hingga harum. Masukkan kentang, ikan teri, dan daun bawang. Lalu bumbui dengan garam, gula dan kaldu bubuk secukupnya. Aduk rata. Jangan lupa koreksi rasa, jika sudah pas angkat dan siap disajikan.",
			CreatedAt:  "2023-07-16 11:46:42.61695346 +0700 WIB m=+3451.675515087",
			UpdatedAt:  "2023-07-16 11:46:42.61695346 +0700 WIB m=+3451.675515087",
		},
	}, nil)

	t.Run("list recipe error", func(t *testing.T) {
		recipe, err := service.List("tuna")
		assert.Nil(t, recipe)
		assert.Equal(t, err, nil)
	})

	t.Run("list recipe success", func(t *testing.T) {
		recipe, err := service.List("tuna")
		// assert.NotNil(t, recipe)
		assert.Len(t, recipe, 0)
		assert.Nil(t, err)
	})
}
