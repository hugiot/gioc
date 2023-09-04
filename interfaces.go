package gioc

type Callback func(sc ServiceContainer) any

type Container interface {
	Get(id string) any
	Has(id string) bool
}

type ServiceContainer interface {
	Container
	Bind(id string, callback Callback)
	Single(id string, callback Callback)
	Instance(id string, instance any)
	Make(id string) any
}

type ServiceProvider interface {
	Register(c ServiceContainer)
	Boot(c ServiceContainer)
}
