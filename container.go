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
	Singleton(key string, f CallbackFunc)
	Make(key string) any
}

func New() Container {
	return &container{
		instances:  make(map[string]any),
		l:          &sync.Mutex{},
		singletons: make(map[string]CallbackFunc),
	}
}

type container struct {
	instances  map[string]any
	providers  []ServiceProvider
	l          *sync.Mutex
	singletons map[string]CallbackFunc
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
	c.l.Lock()
	defer c.l.Unlock()
	if c.hasInstance(key) {
		return
	}
	c.instances[key] = instance
}

func (c *container) Singleton(key string, f CallbackFunc) {
	c.l.Lock()
	defer c.l.Unlock()
	if c.hasSingleton(key) {
		return
	}
	c.singletons[key] = f
}

func (c *container) Bind(key string, f CallbackFunc) {
	c.Instance(key, f)
}

func (c *container) Make(key string) any {
	c.l.Lock()
	defer c.l.Unlock()
	if v, ok := c.instances[key]; ok {
		switch t := v.(type) {
		case CallbackFunc:
			return t()
		default:
			return v
		}
	}

	if c.isSingleton(key) {
		c.instances[key] = c.singletons[key]()
		return c.instances[key]
	}

	return nil
}

func (c *container) hasInstance(key string) bool {
	if _, ok := c.instances[key]; ok {
		return true
	}
	return false
}

func (c *container) hasSingleton(key string) bool {
	if _, ok := c.singletons[key]; ok {
		return true
	}
	return false
}

func (c *container) isSingleton(key string) bool {
	return c.hasSingleton(key)
}
