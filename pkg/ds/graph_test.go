package ds

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	g := NewGraph(7, true)

	g.AddEdge(2, 0, 4)
	g.AddEdge(2, 5, 5)
	g.AddEdge(0, 1, 2)
	g.AddEdge(0, 3, 1)
	g.AddEdge(1, 4, 10)
	g.AddEdge(1, 3, 3)
	g.AddEdge(3, 2, 2)
	g.AddEdge(3, 5, 8)
	g.AddEdge(3, 6, 4)
	g.AddEdge(3, 4, 2)
	g.AddEdge(4, 6, 6)
	g.AddEdge(6, 5, 1)

	log.Printf("%v", g)
}
