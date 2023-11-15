package gioc

type Callback func() any

// Container container
type Container interface {
	// Bind bind multiple instances
	Bind(id string, callback Callback)
	// Single bind singleton
	Single(id string, callback Callback)
	// Instance bind single instance
	Instance(id string, instance any)
	// Make build instance
	Make(id string) any
}

// ServiceProvider service provider
type ServiceProvider interface {
	// Register service register
	Register(c Container)
	// Boot service boot
	Boot(c Container)
}
