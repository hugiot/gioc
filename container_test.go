package gioc

import (
	"fmt"
	"testing"
)

type AppServiceProvider struct {
	c Container
}

func NewAppServiceProvider(c Container) ServiceProvider {
	return &AppServiceProvider{
		c: c,
	}
}

func (a *AppServiceProvider) Register() {
	a.c.Bind("student", func() any {
		return NewStudent()
	})
}

func (a *AppServiceProvider) Boot() {

}

type Student struct {
	ID   string
	Name string
}

func NewStudent() *Student {
	return &Student{
		ID:   "1001",
		Name: "Aim",
	}
}

func TestNew(t *testing.T) {
	c := New()
	c.AddServiceProvider(NewAppServiceProvider(c))
	c.Boot()

	s := c.Make("student").(*Student)
	fmt.Println(s.ID)
	fmt.Println(s.Name)

	s.Name = "edit"

	s2 := c.Make("student").(*Student)
	fmt.Println(s2.ID)
	fmt.Println(s2.Name)
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &Student{
			ID:   "1001",
			Name: "Aim",
		}
	}
}

func BenchmarkIOC(b *testing.B) {
	c := New()
	c.AddServiceProvider(NewAppServiceProvider(c))
	c.Boot()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = c.Make("student").(*Student)
	}
}
