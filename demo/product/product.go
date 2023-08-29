package product

type Product struct {
	Name  string
	Price float64
}

func New(name string, price float64) *Product {
	return &Product{
		Name:  name,
		Price: price,
	}
}
