package database

import (
	"github.com/juliofc/desafio_grpc/internal/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Db *gorm.DB
}

type ProductModel struct {
	ID          int32 `gorm:"primaryKey"`
	Name        string
	Description string
	Price       float64
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{Db: db}
}

func (r *ProductRepository) Save(product *entity.Product) {
	r.Db.Create(&ProductModel{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	})
}

func (r *ProductRepository) FindAll() ([]entity.Product, error) {
	var models []ProductModel
	r.Db.Find(&models)

	var products []entity.Product
	for _, model := range models {
		product := &entity.Product{
			ID:          model.ID,
			Name:        model.Name,
			Description: model.Description,
			Price:       model.Price,
		}
		products = append(products, *product)
	}

	return products, nil
}
