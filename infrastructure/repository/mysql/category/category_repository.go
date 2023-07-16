package category

import (
	"context"
	"errors"
	domain "svc-receipt-luscious/core/domain/category"
	"svc-receipt-luscious/infrastructure/repository/mysql/transactor"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Repository struct {
		db *gorm.DB
	}

	Category struct {
		ID           string `gorm:"primaryKey;column:id"`
		CategoryName string `gorm:"column:category_name"`
		CreatedAt    string `gorm:"column:created_at"`
		UpdatedAt    string `gorm:"column:updated_at"`
	}
)

func (Category) TableName() string {
	return "Categories"
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

func (repo *Repository) GetAllListCategory(categoryName string) ([]domain.Category, error) {
	categories := make([]Category, 0)
	query := repo.db.Table("Categories")

	result := query.Find(&categories)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result.Error = nil
	}

	outData := make([]domain.Category, 0)
	for _, value := range categories {
		var category domain.Category
		category.ID = value.ID
		category.CategoryName = value.CategoryName
		category.CreatedAt = value.CreatedAt
		outData = append(outData, category)
	}

	return outData, nil
}

// func (repo *Repository) GetDetailCategory(categoryID string) (domain.Category, error) {
// 	var category Category

// 	if err := repo.db.First(&category, categoryID).Error; err != nil {
// 		return , err
// 	}

// 	return category, nil
// }

func (repo *Repository) InsertCategory(ctx context.Context, inData *domain.Category) error {
	categories := mappingInput(inData)

	db := repo.getDB(ctx)
	if err := db.Model(categories).Create(&categories).Error; err != nil {
		return err
	}

	return nil
}

func mappingInput(category *domain.Category) Category {
	var result Category

	id := uuid.New()

	timeNow := time.Now()
	result.ID = id.String()
	result.CategoryName = category.CategoryName
	result.CreatedAt = timeNow.String()

	return result
}
