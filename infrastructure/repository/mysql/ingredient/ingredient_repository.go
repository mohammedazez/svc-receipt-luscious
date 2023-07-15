package ingredient

import (
	"context"
	"errors"
	domain "svc-receipt-luscious/core/domain/ingredient"
	"svc-receipt-luscious/infrastructure/repository/mysql/transactor"
	"time"

	"gorm.io/gorm"
)

type (
	Repository struct {
		db *gorm.DB
	}

	Ingredient struct {
		ID             string    `gorm:"primarykey;<-:create"`
		IngredientName string    `gorm:"column:ingredient_name"`
		RecipeID       string    `gorm:"column:recipe_id"`
		Quantity       string    `gorm:"column:quantity"`
		CreatedAt      time.Time `gorm:"column:created_at;<-:create"`
		UpdatedAt      time.Time `gorm:"column:updated_at"`
	}
)

func (Ingredient) TableName() string {
	return "ingredients"
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

func (repo *Repository) GetAllListIngredient(ingredientName string) ([]domain.IngredientService, error) {
	ingredients := make([]Ingredient, 0)
	query := repo.db.Table("Ingredients")

	result := query.Find(&ingredients)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result.Error = nil
	}

	outData := make([]domain.IngredientService, 0)
	for _, value := range ingredients {
		var ingredient domain.IngredientService
		ingredient.ID = value.ID
		ingredient.IngredientName = value.IngredientName
		ingredient.RecipeID = value.RecipeID
		ingredient.Quantity = value.Quantity
		ingredient.CreatedAt = value.CreatedAt
		outData = append(outData, ingredient)
	}

	return outData, nil
}
