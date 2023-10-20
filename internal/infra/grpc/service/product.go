package service

import (
	"context"

	"github.com/juliofc/desafio_grpc/internal/infra/grpc/protofiles/pb"
	"github.com/juliofc/desafio_grpc/internal/usecase"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	CreateProductUseCase   usecase.CreateProductUseCase
	FindAllProductsUseCase usecase.FindAllProductsUseCase
}

func NewProductService(createProductUseCase usecase.CreateProductUseCase, findAllProductsUseCase usecase.FindAllProductsUseCase) *ProductService {
	return &ProductService{
		CreateProductUseCase:   createProductUseCase,
		FindAllProductsUseCase: findAllProductsUseCase,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	dto := usecase.ProductInputDTO{
		Name:        in.Name,
		Description: in.Description,
		Price:       float64(in.Price),
	}
	output := s.CreateProductUseCase.Execute(dto)

	product := &pb.Product{
		Id:          output.ID,
		Name:        output.Name,
		Description: output.Description,
		Price:       float32(output.Price),
	}

	return &pb.CreateProductResponse{
		Product: product,
	}, nil
}

func (s *ProductService) FindProducts(ctx context.Context, in *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	output, err := s.FindAllProductsUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var productsResponse []*pb.Product

	for _, product := range output {
		productResponse := &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       float32(product.Price),
		}

		productsResponse = append(productsResponse, productResponse)
	}

	return &pb.FindProductsResponse{Products: productsResponse}, nil
}
