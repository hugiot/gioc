package main

import (
	"fmt"
	"github.com/hugiot/gioc"
)

const (
	AService string = "service-a"
	BService string = "service-b"
)

type ServiceA struct {
	Server *ServiceB
}

type ServiceB struct {
	Server *ServiceA
}

type AppServiceProvider struct {
}

func (a AppServiceProvider) Register(c gioc.Container) {
	c.Single(BService, func() any {
		aService := c.Make(AService).(*ServiceA)
		return &ServiceB{Server: aService}
	})

	c.Single(AService, func() any {
		bService := c.Make(BService).(*ServiceB)
		return &ServiceA{Server: bService}
	})

}

func (a AppServiceProvider) Boot(c gioc.Container) {
	fmt.Println("boot over")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	gioc.AddServiceProvider(&AppServiceProvider{})
	gioc.Boot()

	//aService := gioc.Make(AService).(*ServiceA)
	//fmt.Println(aService)

	fmt.Println("over")
}
