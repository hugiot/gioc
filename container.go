package gioc

import "sync"

type CallbackFunc func() any

type ServiceProvider interface {
	Register()
	Boot()
}

type Container interface {
	AddServiceProvider(sp ServiceProvider)
	Boot()
	Instance(key string, instance any)
	Bind(key string, f CallbackFunc)
	singleton(key string, f CallbackFunc)
	Make(key string) any
}

func New() Container {
	return &container{
		instances: &sync.Map{},
	}
}

type container struct {
	instances *sync.Map
	providers []ServiceProvider
}

func (c *container) Boot() {
	for i, _ := range c.providers {
		c.providers[i].Register()
	}

	for i, _ := range c.providers {
		c.providers[i].Boot()
	}
}

func (c *container) AddServiceProvider(sp ServiceProvider) {
	c.providers = append(c.providers, sp)
}

func (c *container) Instance(key string, instance any) {
	c.instances.Store(key, instance)
}

func (c *container) singleton(key string, f CallbackFunc) {
	//TODO implement me
	panic("implement me")
}

func (c *container) Bind(key string, f CallbackFunc) {
	c.instances.Store(key, f)
}

func (c *container) Make(key string) any {
	if v, ok := c.instances.Load(key); ok {
		switch tmp := v.(type) {
		case CallbackFunc:
			return tmp()
		default:
			return v
		}
	}

	return nil
}
