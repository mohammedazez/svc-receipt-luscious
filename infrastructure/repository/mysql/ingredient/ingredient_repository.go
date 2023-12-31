package ingredient

import (
	"context"
	"errors"
	domain "svc-receipt-luscious/core/domain/ingredient"
	"svc-receipt-luscious/infrastructure/repository/mysql/transactor"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Repository struct {
		db *gorm.DB
	}

	Ingredient struct {
		ID             string `gorm:"primaryKey;column:id"`
		IngredientName string `gorm:"column:ingredient_name"`
		RecipeID       string `gorm:"column:recipe_id"`
		Quantity       string `gorm:"column:quantity"`
		CreatedAt      string `gorm:"column:created_at"`
		UpdatedAt      string `gorm:"column:updated_at"`
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
	query := repo.db.Table("ingredients")

	if ingredientName != "" {
		query = query.Where("ingredient_name LIKE ?", "%"+ingredientName+"%")
	}

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
		ingredient.UpdatedAt = value.UpdatedAt
		outData = append(outData, ingredient)
	}

	return outData, nil
}

func (repo *Repository) InsertIngredient(ctx context.Context, inData *domain.IngredientService) error {
	ingredients := mappingInput(inData)

	db := repo.getDB(ctx)
	if err := db.Model(ingredients).Create(&ingredients).Error; err != nil {
		return err
	}

	return nil
}

func (repo *Repository) UpdateIngredient(ctx context.Context, inData *domain.IngredientService) error {
	ingredients := mappingInput(inData)

	db := repo.getDB(ctx)
	timeNow := time.Now()
	ingredients.UpdatedAt = timeNow.String()

	updates := map[string]interface{}{
		"ingredient_name": ingredients.IngredientName,
		"recipe_id":       ingredients.RecipeID,
		"quantity":        ingredients.Quantity,
		"updated_at":      ingredients.UpdatedAt,
	}

	err := db.Model(&Ingredient{}).Where("id = ?", inData.ID).Updates(updates).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) DeleteIngredient(ctx context.Context, ingredientID string) error {
	db := repo.getDB(ctx)
	err := db.Where("id = ?", ingredientID).Delete(&Ingredient{}).Error
	if err != nil {
		return err
	}

	return nil
}

func mappingInput(ingredient *domain.IngredientService) Ingredient {
	var result Ingredient

	id := uuid.New()

	timeNow := time.Now()
	result.ID = id.String()
	result.IngredientName = ingredient.IngredientName
	result.RecipeID = ingredient.RecipeID
	result.Quantity = ingredient.Quantity
	result.CreatedAt = timeNow.String()

	return result
}
