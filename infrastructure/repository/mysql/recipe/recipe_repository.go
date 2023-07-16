package recipe

import (
	"context"
	"errors"
	"svc-receipt-luscious/infrastructure/repository/mysql/transactor"

	domain "svc-receipt-luscious/core/domain/recipe"

	"gorm.io/gorm"
)

type (
	Repository struct {
		db *gorm.DB
	}

	Recipe struct {
		ID         string `gorm:"primaryKey;column:id"`
		RecipeName string `gorm:"column:recipe_name"`
		CreatedAt  string `gorm:"column:created_at"`
		UpdatedAt  string `gorm:"column:updated_at"`
	}
)

func (Recipe) TableName() string {
	return "Recipes"
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) getDB(ctx context.Context) *gorm.DB {
	dbWithTx := transactor.ExtractTx(ctx)
	if dbWithTx != nil {
		return dbWithTx
	}
	return repo.db
}

func (repo *Repository) GetAllListRecipe(recipeName string) ([]domain.Recipe, error) {
	recipes := make([]Recipe, 0)
	query := repo.db.Table("Recipes")

	result := query.Find(&recipes)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result.Error = nil
	}

	outData := make([]domain.Recipe, 0)
	for _, value := range recipes {
		var recipe domain.Recipe
		recipe.ID = value.ID
		recipe.RecipeName = value.RecipeName
		recipe.UpdatedAt = value.UpdatedAt
		recipe.CreatedAt = value.CreatedAt
		outData = append(outData, recipe)
	}

	return outData, nil
}
