package recipe

import (
	"context"
	"errors"
	"svc-receipt-luscious/infrastructure/repository/mysql/transactor"
	"time"

	domain "svc-receipt-luscious/core/domain/recipe"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Repository struct {
		db *gorm.DB
	}

	Recipe struct {
		ID         string `gorm:"primaryKey;column:id"`
		RecipeName string `gorm:"column:recipe_name"`
		CategoryID string `gorm:"column:category_id"`
		HowToMake  string `gorm:"column:how_to_make"`
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

func (repo *Repository) InsertRecipe(ctx context.Context, inData *domain.Recipe) error {
	recipes := mappingInput(inData)

	db := repo.getDB(ctx)
	if err := db.Model(recipes).Create(&recipes).Error; err != nil {
		return err
	}

	return nil
}

func mappingInput(recipe *domain.Recipe) Recipe {
	var result Recipe

	id := uuid.New()

	timeNow := time.Now()
	result.ID = id.String()
	result.RecipeName = recipe.RecipeName
	result.CategoryID = recipe.CategoryID
	result.HowToMake = recipe.HowToMake
	result.CreatedAt = timeNow.String()

	return result
}
