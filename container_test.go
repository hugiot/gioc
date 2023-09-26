package gioc

import (
	"log"
	"os"
	"testing"
)

type appServiceProvider struct {
}

func (*appServiceProvider) Register(c Container) {
	c.Bind("log", func() any {
		return log.New(os.Stderr, "custom | ", log.LstdFlags)
	})

	c.Single("single-log", func() any {
		return log.New(os.Stderr, "custom | ", log.LstdFlags)
	})

	instance := log.New(os.Stderr, "custom | ", log.LstdFlags)
	c.Instance("instance-log", instance)
}

func (*appServiceProvider) Boot(c Container) {
}

func initContainer() {
	AddServerProvider(&appServiceProvider{})
	Boot()
}

func BenchmarkBind(b *testing.B) {
	initContainer()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Make("log").(*log.Logger)
	}
}

func BenchmarkSingle(b *testing.B) {
	initContainer()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Make("single-log").(*log.Logger)
	}
}

func BenchmarkInstance(b *testing.B) {
	initContainer()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Make("instance-log").(*log.Logger)
	}
}
