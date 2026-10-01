package repository

import (
	"context"
	"errors"

	"github.com/faizallmaullana/rekrutment-gbu-go/internal/domain"
	"github.com/faizallmaullana/rekrutment-gbu-go/internal/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	database *gorm.DB
}

func NewProductRepository(database *gorm.DB) *ProductRepository {
	return &ProductRepository{database: database}
}

func (repository *ProductRepository) List(ctx context.Context) ([]domain.Product, error) {
	var models []models.ProductModel
	if err := repository.database.WithContext(ctx).Order("id ASC").Find(&models).Error; err != nil {
		return nil, err
	}

	products := make([]domain.Product, 0, len(models))
	for _, model := range models {
		products = append(products, model.ToDomain())
	}

	return products, nil
}

func (repository *ProductRepository) Get(ctx context.Context, id int64) (domain.Product, error) {
	var model models.ProductModel
	err := repository.database.WithContext(ctx).First(&model, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.Product{}, domain.ErrNotFound
	}
	if err != nil {
		return domain.Product{}, err
	}

	return model.ToDomain(), nil
}

func (repository *ProductRepository) Create(ctx context.Context, product domain.Product) (domain.Product, error) {
	model := models.ProductModelFromDomain(product)
	if err := repository.database.WithContext(ctx).Create(&model).Error; err != nil {
		return domain.Product{}, err
	}

	return model.ToDomain(), nil
}

func (repository *ProductRepository) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	result := repository.database.WithContext(ctx).
		Model(&models.ProductModel{}).
		Where("id = ?", product.ID).
		Updates(map[string]any{"name": product.Name, "price": product.Price})
	if result.Error != nil {
		return domain.Product{}, result.Error
	}
	if result.RowsAffected == 0 {
		return domain.Product{}, domain.ErrNotFound
	}

	return repository.Get(ctx, product.ID)
}

func (repository *ProductRepository) Delete(ctx context.Context, id int64) error {
	result := repository.database.WithContext(ctx).Delete(&models.ProductModel{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}
