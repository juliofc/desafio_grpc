package entity

type Product struct {
	ID          int32
	Name        string
	Description string
	Price       float64
}

func NewProduct(name string, description string, price float64) *Product {
	return &Product{
		ID:          0,
		Name:        name,
		Description: description,
		Price:       price,
	}
}
