package gioc

import (
	"sync"
)

// c default ServiceContainer instance
var c = createContainer()

// data storage struct
type data struct {
	once      *sync.Once
	singleton bool
	callback  Callback // bind and singleton
	instance  any      // singleton and instance
}

// getSingleton get singleton
func (sd *data) getSingleton() any {
	sd.once.Do(func() {
		sd.instance = sd.callback()
	})
	return sd.instance
}

// getInstanceOrBind get instance or bind
func (sd *data) getInstanceOrBind() any {
	if sd.instance != nil {
		return sd.instance
	}
	return sd.callback()
}

// Make generate instance
func (sd *data) Make() any {
	if sd.singleton {
		return sd.getSingleton()
	}
	return sd.getInstanceOrBind()
}

// container implementation ServiceContainer
type container struct {
	m         *sync.Map
	providers []ServiceProvider
	boot      bool
	lock      *sync.Mutex
}

// createContainer create Container instance
func createContainer() *container {
	return &container{
		m:         &sync.Map{},
		providers: make([]ServiceProvider, 0),
		lock:      &sync.Mutex{},
	}
}

func (c *container) Bind(id string, callback Callback) {
	if _, ok := c.m.Load(id); !ok {
		c.m.Store(id, &data{callback: callback})
	}
}

func (c *container) Single(id string, callback Callback) {
	if _, ok := c.m.Load(id); !ok {
		c.m.Store(id, &data{once: &sync.Once{}, singleton: true, callback: callback})
	}
}

func (c *container) Instance(id string, instance any) {
	if _, ok := c.m.Load(id); !ok {
		c.m.Store(id, &data{instance: instance})
	}
}

func (c *container) Make(id string) any {
	if v, ok := c.m.Load(id); ok {
		sd := v.(*data)
		return sd.Make()
	}
	return nil
}

func (c *container) AddServiceProvider(sp ServiceProvider) {
	c.providers = append(c.providers, sp)
}

func (c *container) Boot() {
	if len(c.providers) > 0 {
		for i, _ := range c.providers {
			c.providers[i].Register(c)
		}

		for i, _ := range c.providers {
			c.providers[i].Boot(c)
		}

		// preload
		c.m.Range(func(key, value any) bool {
			_ = c.Make(key.(string))
			return true
		})
	}
}

func AddServiceProvider(sp ServiceProvider) {
	c.AddServiceProvider(sp)
}

func Make(id string) any {
	return c.Make(id)
}

func Boot() {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.boot == false {
		c.boot = true
		c.Boot()
	}
}
