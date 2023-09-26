package gioc

import (
	"sync"
)

// c default ServiceContainer instance
var c = New()

// data storage struct
type data struct {
	once      *sync.Once
	singleton bool
	callback  Callback // bind and singleton
	instance  any      // singleton and instance
}

// GetSingleton get singleton
func (sd *data) GetSingleton() any {
	sd.once.Do(func() {
		sd.instance = sd.callback()
	})
	return sd.instance
}

// GetInstanceOrBind get instance or bind
func (sd *data) GetInstanceOrBind() any {
	if sd.instance != nil {
		return sd.instance
	}
	return sd.callback()
}

// Make generate instance
func (sd *data) Make() any {
	if sd.singleton {
		return sd.GetSingleton()
	}
	return sd.GetInstanceOrBind()
}

// sc implementation ServiceContainer
type sc struct {
	m         *sync.Map
	providers []ServiceProvider
}

// New Container instance
func New() ServiceContainer {
	return &sc{
		m:         &sync.Map{},
		providers: make([]ServiceProvider, 0),
	}
}

func (c *sc) Bind(id string, callback Callback) {
	if _, ok := c.m.Load(id); !ok {
		c.m.Store(id, &data{callback: callback})
	}
}

func (c *sc) Single(id string, callback Callback) {
	if _, ok := c.m.Load(id); !ok {
		c.m.Store(id, &data{once: &sync.Once{}, singleton: true, callback: callback})
	}
}

func (c *sc) Instance(id string, instance any) {
	if _, ok := c.m.Load(id); !ok {
		c.m.Store(id, &data{instance: instance})
	}
}

func (c *sc) Make(id string) any {
	if v, ok := c.m.Load(id); ok {
		sd := v.(*data)
		return sd.Make()
	}
	return nil
}

func (c *sc) AddServerProvider(sp ServiceProvider) {
	c.providers = append(c.providers, sp)
}

func (c *sc) Boot() {
	if len(c.providers) > 0 {
		for i, _ := range c.providers {
			c.providers[i].Register(c)
		}

		for i, _ := range c.providers {
			c.providers[i].Boot(c)
		}
	}
}

func AddServerProvider(sp ServiceProvider) {
	c.AddServerProvider(sp)
}

func Make(id string) any {
	return c.Make(id)
}

func Boot() {
	c.Boot()
}
