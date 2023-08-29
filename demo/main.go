package main

import (
	"fmt"
	"github.com/hugiot/gioc/demo/order"
	"github.com/hugiot/gioc/demo/provider"
	"github.com/hugiot/gioc/src/container"
)

func main() {
	container.AddServerProvider(&provider.AppServiceProvider{})
	container.Boot()

	// 简单绑定
	o := container.Make(provider.OrderService).(*order.Order)
	fmt.Println(o.ID, o.Product.Name, o.Product.Price)
	o.Product.Name = "edit"

	o2 := container.Make(provider.OrderService).(*order.Order)
	fmt.Println(o2.ID, o2.Product.Name, o2.Product.Price)

	// 单例
	o3 := container.Make(provider.OrderSingleService).(*order.Order)
	fmt.Println(o3.ID, o3.Product.Name, o3.Product.Price)
	o3.Product.Name = "edit"

	o4 := container.Make(provider.OrderSingleService).(*order.Order)
	fmt.Println(o4.ID, o4.Product.Name, o4.Product.Price)
}
