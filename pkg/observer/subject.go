package observer

type Subject interface {
	Register(Observer)
	Deregister(string)
	NotifyAll()
}

func NewItem(s string) *item {
	return &item{
		name: s,
	}
}

type item struct {
	observers []Observer
	name      string
}

func (it *item) Register(o Observer) {
	it.observers = append(it.observers, o)
}

func (it *item) Deregister(id string) {
	for i := range it.observers {
		if it.observers[i].ID() == id {
			it.observers[i] = it.observers[len(it.observers)-1]
			it.observers = it.observers[:len(it.observers)-1]
		}
	}
}

func (it *item) NotifyAll() {
	for i := range it.observers {
		it.observers[i].Update(it.name)
	}
}
