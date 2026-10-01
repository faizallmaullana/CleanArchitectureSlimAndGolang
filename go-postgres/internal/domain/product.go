package domain

import (
	"context"
	"errors"
	"strings"
)

var (
	ErrNotFound = errors.New("product not found")
	ErrInvalid  = errors.New("invalid product")
)

type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewProduct(name string, price float64) (Product, error) {
	product := Product{Name: strings.TrimSpace(name), Price: price}
	if err := product.Validate(); err != nil {
		return Product{}, err
	}

	return product, nil
}

func (product Product) Validate() error {
	if product.Name == "" {
		return errors.Join(ErrInvalid, errors.New("product name is required"))
	}
	if product.Price < 0 {
		return errors.Join(ErrInvalid, errors.New("product price must be zero or greater"))
	}

	return nil
}

func (product *Product) Update(name string, price float64) error {
	updated, err := NewProduct(name, price)
	if err != nil {
		return err
	}

	product.Name = updated.Name
	product.Price = updated.Price
	return nil
}

type ProductRepository interface {
	List(context.Context) ([]Product, error)
	Get(context.Context, int64) (Product, error)
	Create(context.Context, Product) (Product, error)
	Update(context.Context, Product) (Product, error)
	Delete(context.Context, int64) error
}
