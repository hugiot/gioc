package gioc

import (
	"fmt"
	"sync"
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
	var s *Student
	once := &sync.Once{}
	a.c.Bind("student", func() any {
		once.Do(func() {
			s = &Student{
				ID:   "1001",
				Name: "Aim",
			}
		})
		return s
	})
}

func (a *AppServiceProvider) Boot() {

}

type Student struct {
	ID   string
	Name string
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
