package models

import "github.com/faizallmaullana/rekrutment-gbu-go/internal/domain"

type ProductModel struct {
	ID    int64   `gorm:"primaryKey"`
	Name  string  `gorm:"column:name"`
	Price float64 `gorm:"column:price"`
}

func (ProductModel) TableName() string {
	return "products"
}

func (model ProductModel) ToDomain() domain.Product {
	return domain.Product{ID: model.ID, Name: model.Name, Price: model.Price}
}

func ProductModelFromDomain(product domain.Product) ProductModel {
	return ProductModel{ID: product.ID, Name: product.Name, Price: product.Price}
}
