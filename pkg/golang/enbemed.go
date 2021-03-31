package golang

import "fmt"

type Man interface {
	Walk()
	Talk()
	Play()
	Shop()
}

type PlayerGetter interface {
	Play()
}

type walker interface {
	Walk()
}

var _ interface{} = struct{}{}

type father struct{}

func (f *father) Walk() { fmt.Println("father walk") }
func (f *father) Talk() { fmt.Println("father talk") }
func (f *father) Play() { fmt.Println("father Play") }
func (f *father) Shop() { fmt.Println("father Shop") }

type son struct {
	*father
}

func (f *son) Play() { fmt.Println("son Play") }
func (f *son) Shop() { fmt.Println("son Shop") }

func NewMan() Man {
	return &son{
		father: &father{},
	}
}
