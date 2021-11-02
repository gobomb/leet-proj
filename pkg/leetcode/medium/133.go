package medium

/*
	133. Clone Graph
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

// rename symbol to Node
type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

type GraphNodeMaker struct {
	list  []*GraphNode
	input [][]int
}

func makeGraphNode(input [][]int) *GraphNode {
	m := &GraphNodeMaker{
		list:  make([]*GraphNode, len(input)),
		input: input,
	}
	if len(m.list) == 0 {
		return nil
	}
	return m.makeGraphNode(1)
}

func (m GraphNodeMaker) makeGraphNode(val int) *GraphNode {
	g := &GraphNode{
		Val: val,
	}
	m.list[val-1] = g
	for _, neighbor := range m.input[val-1] {
		if m.list[neighbor-1] == nil {
			g.Neighbors = append(g.Neighbors, m.makeGraphNode(neighbor))
		} else {
			g.Neighbors = append(g.Neighbors, m.list[neighbor-1])
		}
	}
	return g
}

func cloneGraph(node *GraphNode) *GraphNode {
	return graphCloner{
		make(map[int]*GraphNode),
	}.cloneGraph(node)
}

type graphCloner struct {
	visited map[int]*GraphNode
}

func (gc graphCloner) cloneGraph(node *GraphNode) *GraphNode {
	if node == nil {
		return nil
	}
	if v, ok := gc.visited[node.Val]; ok {
		return v
	}

	nn := &GraphNode{Val: node.Val}
	gc.visited[nn.Val] = nn

	for _, v := range node.Neighbors {
		nn.Neighbors = append(nn.Neighbors, gc.cloneGraph(v))
	}

	return nn
}
