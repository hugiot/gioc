package interfaces

type ContainerCallback func(sc ServiceContainer) any

type Container interface {
	Get(id string) any
	Has(id string) bool
}

type ServiceContainer interface {
	Container
	Bind(id string, callback ContainerCallback)
	Single(id string, callback ContainerCallback)
	Instance(id string, instance any)
	Make(id string) any
}

type ServiceProvider interface {
	Register(c ServiceContainer)
	Boot(c ServiceContainer)
}
