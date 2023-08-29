package order

import (
	"github.com/hugiot/gioc/demo/product"
	"strconv"
	"time"
)

type Order struct {
	ID        string
	Product   *product.Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(p *product.Product) *Order {
	return &Order{
		ID:        strconv.Itoa(int(time.Now().UnixNano())),
		Product:   p,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
