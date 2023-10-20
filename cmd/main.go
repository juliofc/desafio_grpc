package main

import (
	"fmt"
	"net"

	"github.com/juliofc/desafio_grpc/internal/infra/database"
	"github.com/juliofc/desafio_grpc/internal/infra/grpc/protofiles/pb"
	"github.com/juliofc/desafio_grpc/internal/infra/grpc/service"
	"github.com/juliofc/desafio_grpc/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	dsn := "./database.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	defer db.DB()

	db.AutoMigrate(&database.ProductModel{})

	productRepository := database.NewProductRepository(db)
	createProductUseCase := usecase.NewCreateProductUseCase(productRepository)
	findAllProductsUseCase := usecase.NewFindAllProductsUseCase(productRepository)

	productService := service.NewProductService(*createProductUseCase, *findAllProductsUseCase)

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, productService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port 50051")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
