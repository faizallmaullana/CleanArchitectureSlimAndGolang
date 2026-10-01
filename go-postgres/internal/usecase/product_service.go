package usecase

import (
	"context"

	"github.com/example/rekrutment-gbu-go/internal/domain"
)

type ProductService struct {
	repository domain.ProductRepository
}

func NewProductService(repository domain.ProductRepository) *ProductService {
	return &ProductService{repository: repository}
}

func (service *ProductService) List(ctx context.Context) ([]domain.Product, error) {
	return service.repository.List(ctx)
}

func (service *ProductService) Get(ctx context.Context, id int64) (domain.Product, error) {
	return service.repository.Get(ctx, id)
}

func (service *ProductService) Create(ctx context.Context, name string, price float64) (domain.Product, error) {
	product, err := domain.NewProduct(name, price)
	if err != nil {
		return domain.Product{}, err
	}

	return service.repository.Create(ctx, product)
}

func (service *ProductService) Update(ctx context.Context, id int64, name string, price float64) (domain.Product, error) {
	product, err := service.repository.Get(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}

	if err := product.Update(name, price); err != nil {
		return domain.Product{}, err
	}

	return service.repository.Update(ctx, product)
}

func (service *ProductService) Delete(ctx context.Context, id int64) error {
	return service.repository.Delete(ctx, id)
}