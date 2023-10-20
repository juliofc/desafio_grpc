package usecase

import (
	"github.com/juliofc/desafio_grpc/internal/entity"
)

type FindAllProductsOutputDTO struct {
	ID          int32
	Name        string
	Description string
	Price       float64
}

type FindAllProductsUseCase struct {
	ProductRepository entity.ProductRepositoryInterface
}

func NewFindAllProductsUseCase(
	productRepository entity.ProductRepositoryInterface,
) *FindAllProductsUseCase {
	return &FindAllProductsUseCase{
		ProductRepository: productRepository,
	}
}

func (f *FindAllProductsUseCase) Execute() ([]FindAllProductsOutputDTO, error) {
	products, err := f.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var list []FindAllProductsOutputDTO

	for _, product := range products {
		dto := FindAllProductsOutputDTO{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		}

		list = append(list, dto)
	}

	return list, nil
}
