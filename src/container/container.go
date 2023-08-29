package container

import (
	"github.com/hugiot/gioc/src/interfaces"
	"sync"
)

var c = New()

type storeData struct {
	singleton bool
	callback  func(sc interfaces.ServiceContainer) any
	instance  any
}

type Container struct {
	m         *sync.Map
	providers []interfaces.ServiceProvider
}

func New() *Container {
	return &Container{
		m:         &sync.Map{},
		providers: make([]interfaces.ServiceProvider, 0),
	}
}

func (c *Container) Get(id string) any {
	return c.Make(id)
}

func (c *Container) Has(id string) bool {
	_, ok := c.m.Load(id)
	return ok
}

func (c *Container) Bind(id string, callback func(sc interfaces.ServiceContainer) any) {
	if !c.Has(id) {
		c.m.Store(id, &storeData{
			callback: callback,
		})
	}
}

func (c *Container) Single(id string, callback func(sc interfaces.ServiceContainer) any) {
	if !c.Has(id) {
		c.m.Store(id, &storeData{
			singleton: true,
			callback:  callback,
		})
	}
}

func (c *Container) Instance(id string, instance any) {
	if !c.Has(id) {
		c.m.Store(id, &storeData{
			singleton: true,
			instance:  instance,
		})
	}
}

func (c *Container) Make(id string) any {
	if v, ok := c.m.Load(id); ok {
		sd := v.(*storeData)
		if sd.singleton {
			if sd.instance == nil {
				sd.instance = sd.callback(c)
			}
			return sd.instance
		} else {
			return sd.callback(c)
		}
	}
	return nil
}

func (c *Container) AddServerProvider(sp interfaces.ServiceProvider) {
	c.providers = append(c.providers, sp)
}

func (c *Container) Boot() {
	if len(c.providers) > 0 {
		for i, _ := range c.providers {
			c.providers[i].Register(c)
		}

		for i, _ := range c.providers {
			c.providers[i].Boot(c)
		}
	}
}

func Get(id string) any {
	return c.Get(id)
}

func Has(id string) bool {
	return c.Has(id)
}

func AddServerProvider(sp interfaces.ServiceProvider) {
	c.AddServerProvider(sp)
}

func Bind(id string, callback func(sc interfaces.ServiceContainer) any) {
	c.Bind(id, callback)
}

func Single(id string, callback func(sc interfaces.ServiceContainer) any) {
	c.Single(id, callback)
}

func Instance(id string, instance any) {
	c.Instance(id, instance)
}

func Make(id string) any {
	return c.Make(id)
}

func Boot() {
	c.Boot()
}
