package golang

import "fmt"

type Duck struct {
	height int
	name   string
}

type Option func(*Duck)

func WithHeight(height int) Option {
	return func(d *Duck) {
		d.height = height
	}
}

func WithName(name string) Option {
	return func(d *Duck) {
		d.name = name
	}
}

// 《Go语言精进之路》25.4·
func NewDuck(opts ...Option) *Duck {
	d := &Duck{
		height: 1,
		name:   "default",
	}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

func (d *Duck) Print() {
	fmt.Println(d.name, d.height)
}
