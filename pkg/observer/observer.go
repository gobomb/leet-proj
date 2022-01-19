package observer

import "log"

type Observer interface {
	Update(string)
	ID() string
}

type customer struct {
	id string
}

func (c *customer) Update(s string) {
	log.Printf("%v updated by %v", c.ID(), s)
}

func (c *customer) ID() string {
	return c.id
}
