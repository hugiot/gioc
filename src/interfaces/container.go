package interfaces

type Container interface {
	Get(id string) any
	Has(id string) bool
}

type ServiceContainer interface {
	Container
	Bind(id string, callback func(sc ServiceContainer) any)
	Single(id string, callback func(sc ServiceContainer) any)
	Instance(id string, instance any)
	Make(id string) any
}

type ServiceProvider interface {
	Register(c ServiceContainer)
	Boot(c ServiceContainer)
}
