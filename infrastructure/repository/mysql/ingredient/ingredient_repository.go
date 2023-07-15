package ingredient

import (
	"context"
	domain "svc-receipt-luscious/core/domain/ingredient"
	"svc-receipt-luscious/infrastructure/repository/mysql/transactor"

	"gorm.io/gorm"
)

type (
	Repository struct {
		db *gorm.DB
	}

	Ingredient struct {
		ID string `gorm:"primarykey;<-:create"`
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

}
