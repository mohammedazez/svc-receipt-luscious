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
	return "categories"
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

	query := repo.db.Table("categories")

	if categoryName != "" {
		query = query.Where("category_name LIKE ?", "%"+categoryName+"%")
	}

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
		category.UpdatedAt = value.UpdatedAt
		outData = append(outData, category)
	}

	return outData, nil
}

func (repo *Repository) GetDetailCategory(categoryID string) (*domain.Category, error) {
	var Categories domain.Category

	result := repo.db.Joins("JOIN recipes ON categories.id = recipes.category_id").
		Select("categories.*, recipes.*").
		Where("categories.id = ?", categoryID).
		First(&Categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return &Categories, nil
}

func (repo *Repository) InsertCategory(ctx context.Context, inData *domain.Category) error {
	categories := mappingInput(inData)

	db := repo.getDB(ctx)
	if err := db.Model(categories).Create(&categories).Error; err != nil {
		return err
	}

	return nil
}

func (repo *Repository) UpdateCategory(ctx context.Context, inData *domain.Category) error {
	categories := mappingInput(inData)

	db := repo.getDB(ctx)
	timeNow := time.Now()
	categories.UpdatedAt = timeNow.String()

	updates := map[string]interface{}{
		"category_name": categories.CategoryName,
		"updated_at":    categories.UpdatedAt,
	}

	err := db.Model(&Category{}).Where("id = ?", inData.ID).Updates(updates).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) DeleteCategory(ctx context.Context, categoryID string) error {
	db := repo.getDB(ctx)
	err := db.Where("id = ?", categoryID).Delete(&Category{}).Error
	if err != nil {
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
