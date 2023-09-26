package main

import (
	"github.com/hugiot/gioc"
	"log"
	"os"
)

const (
	LogService string = "log"
)

type AppServiceProvider struct {
}

func (a AppServiceProvider) Register(c gioc.Container) {
	c.Single(LogService, func() any {
		return log.New(os.Stderr, "custom | ", log.LstdFlags)
	})
}

func (a AppServiceProvider) Boot(c gioc.Container) {
}

func main() {
	gioc.AddServerProvider(&AppServiceProvider{})
	gioc.Boot()
	logger := gioc.Make(LogService).(*log.Logger)
	logger.Println("this is content")
	logger.SetPrefix("edit | ")
	logger.Println("this is content")

	logger2 := gioc.Make(LogService).(*log.Logger)
	logger2.Println("this is content")
}
