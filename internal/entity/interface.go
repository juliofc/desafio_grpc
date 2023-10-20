package entity

type ProductRepositoryInterface interface {
	Save(product *Product)
	FindAll() ([]Product, error)
}
