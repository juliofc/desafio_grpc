package usecase

import (
	"github.com/juliofc/desafio_grpc/internal/entity"
)

type ProductInputDTO struct {
	Name        string
	Description string
	Price       float64
}

type CreateProductOutputDTO struct {
	ID          int32
	Name        string
	Description string
	Price       float64
}

type CreateProductUseCase struct {
	ProductRepository entity.ProductRepositoryInterface
}

func NewCreateProductUseCase(
	productRepository entity.ProductRepositoryInterface,
) *CreateProductUseCase {
	return &CreateProductUseCase{
		ProductRepository: productRepository,
	}
}

func (c *CreateProductUseCase) Execute(input ProductInputDTO) CreateProductOutputDTO {
	product := entity.NewProduct(input.Name, input.Description, input.Price)

	c.ProductRepository.Save(product)

	dto := CreateProductOutputDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}

	return dto
}
