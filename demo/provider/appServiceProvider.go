package provider

import (
	"github.com/hugiot/gioc/demo/order"
	"github.com/hugiot/gioc/demo/product"
	"github.com/hugiot/gioc/src/interfaces"
)

const (
	ProductService string = "product"
	OrderService   string = "order"

	ProductSingleService string = "product-single"
	OrderSingleService   string = "order-single"
)

type AppServiceProvider struct {
}

func (sp *AppServiceProvider) Register(c interfaces.ServiceContainer) {
	c.Bind(ProductService, func(sc interfaces.ServiceContainer) any {
		return product.New("pc", 100)
	})
	c.Bind(OrderService, func(sc interfaces.ServiceContainer) any {
		p := sc.Make(ProductService).(*product.Product)
		return order.New(p)
	})

	c.Single(ProductSingleService, func(sc interfaces.ServiceContainer) any {
		return product.New("phone", 999)
	})
	c.Single(OrderSingleService, func(sc interfaces.ServiceContainer) any {
		p := sc.Make(ProductSingleService).(*product.Product)
		return order.New(p)
	})
}

func (sp *AppServiceProvider) Boot(c interfaces.ServiceContainer) {

}
