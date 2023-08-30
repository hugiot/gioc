package container

import (
	"github.com/hugiot/gioc/src/interfaces"
	"sync"
)

// c default Container instance
var c = New()

// data storage struct
type data struct {
	once      *sync.Once
	singleton bool
	callback  interfaces.ContainerCallback // bind and singleton
	instance  any                          // singleton and instance
}

// GetSingleton get singleton
func (sd *data) GetSingleton(c interfaces.ServiceContainer) any {
	sd.once.Do(func() {
		sd.instance = sd.callback(c)
	})
	return sd.instance
}

// GetInstanceOrBind get instance or bind
func (sd *data) GetInstanceOrBind(c interfaces.ServiceContainer) any {
	if sd.instance != nil {
		return sd.instance
	}
	return sd.callback(c)
}

// Make generate instance
func (sd *data) Make(c interfaces.ServiceContainer) any {
	if sd.singleton {
		return sd.GetSingleton(c)
	}
	return sd.GetInstanceOrBind(c)
}

// Container implementation interfaces.ServiceContainer
type Container struct {
	m         *sync.Map
	providers []interfaces.ServiceProvider
}

// New Container instance
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

func (c *Container) Bind(id string, callback interfaces.ContainerCallback) {
	if !c.Has(id) {
		c.m.Store(id, c.generateBindData(callback))
	}
}

func (c *Container) Single(id string, callback interfaces.ContainerCallback) {
	if !c.Has(id) {
		c.m.Store(id, c.generateSingletonData(callback))
	}
}

func (c *Container) Instance(id string, instance any) {
	if !c.Has(id) {
		c.m.Store(id, c.generateInstanceData(instance))
	}
}

func (c *Container) Make(id string) any {
	if v, ok := c.m.Load(id); ok {
		sd := v.(*data)
		return sd.Make(c)
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

func (c *Container) generateBindData(callback interfaces.ContainerCallback) *data {
	return &data{callback: callback}
}

func (c *Container) generateSingletonData(callback interfaces.ContainerCallback) *data {
	return &data{once: &sync.Once{}, singleton: true, callback: callback}
}

func (c *Container) generateInstanceData(instance any) *data {
	return &data{instance: instance}
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

func Bind(id string, callback interfaces.ContainerCallback) {
	c.Bind(id, callback)
}

func Single(id string, callback interfaces.ContainerCallback) {
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
