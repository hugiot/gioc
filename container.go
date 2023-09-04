package gioc

import (
	"sync"
)

// c default Container instance
var c = New()

// data storage struct
type data struct {
	once      *sync.Once
	singleton bool
	callback  Callback // bind and singleton
	instance  any      // singleton and instance
}

// GetSingleton get singleton
func (sd *data) GetSingleton(c ServiceContainer) any {
	sd.once.Do(func() {
		sd.instance = sd.callback(c)
	})
	return sd.instance
}

// GetInstanceOrBind get instance or bind
func (sd *data) GetInstanceOrBind(c ServiceContainer) any {
	if sd.instance != nil {
		return sd.instance
	}
	return sd.callback(c)
}

// Make generate instance
func (sd *data) Make(c ServiceContainer) any {
	if sd.singleton {
		return sd.GetSingleton(c)
	}
	return sd.GetInstanceOrBind(c)
}

// C implementation ServiceContainer
type C struct {
	m         *sync.Map
	providers []ServiceProvider
}

// New Container instance
func New() *C {
	return &C{
		m:         &sync.Map{},
		providers: make([]ServiceProvider, 0),
	}
}

func (c *C) Get(id string) any {
	return c.Make(id)
}

func (c *C) Has(id string) bool {
	_, ok := c.m.Load(id)
	return ok
}

func (c *C) Bind(id string, callback Callback) {
	if !c.Has(id) {
		c.m.Store(id, c.generateBindData(callback))
	}
}

func (c *C) Single(id string, callback Callback) {
	if !c.Has(id) {
		c.m.Store(id, c.generateSingletonData(callback))
	}
}

func (c *C) Instance(id string, instance any) {
	if !c.Has(id) {
		c.m.Store(id, c.generateInstanceData(instance))
	}
}

func (c *C) Make(id string) any {
	if v, ok := c.m.Load(id); ok {
		sd := v.(*data)
		return sd.Make(c)
	}
	return nil
}

func (c *C) AddServerProvider(sp ServiceProvider) {
	c.providers = append(c.providers, sp)
}

func (c *C) Boot() {
	if len(c.providers) > 0 {
		for i, _ := range c.providers {
			c.providers[i].Register(c)
		}

		for i, _ := range c.providers {
			c.providers[i].Boot(c)
		}
	}
}

func (c *C) generateBindData(callback Callback) *data {
	return &data{callback: callback}
}

func (c *C) generateSingletonData(callback Callback) *data {
	return &data{once: &sync.Once{}, singleton: true, callback: callback}
}

func (c *C) generateInstanceData(instance any) *data {
	return &data{instance: instance}
}

func Get(id string) any {
	return c.Get(id)
}

func Has(id string) bool {
	return c.Has(id)
}

func AddServerProvider(sp ServiceProvider) {
	c.AddServerProvider(sp)
}

func Bind(id string, callback Callback) {
	c.Bind(id, callback)
}

func Single(id string, callback Callback) {
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
