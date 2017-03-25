// implementation of BFS and DFS graph traversing methods
package graph

type Node struct {
	Value   interface{}
	Visited bool
	Edges   []*Edge
}

type Edge struct {
	Value    interface{}
	NodeFrom *Node
	NodeTo   *Node
}

type Graph struct {
	Nodes     []*Node
	Edges     []*Edge
	NodeNames []string
	nodeMap   map[*Node]bool
}

func NewGraph() *Graph {
	return &Graph{Nodes: []*Node{}, Edges: []*Edge{}, NodeNames: []string{}, nodeMap: map[*Node]bool{}}
}

func (r *Graph) SetNodeNames(nodeNames []string) {
	r.NodeNames = nodeNames
}

func (r *Graph) InsertNode(newNodeVal interface{}) *Node {
	node := Node{Value: newNodeVal}
	r.Nodes = append(r.Nodes, &node)
	r.nodeMap[&node] = true
	return &node
}

func (r *Graph) InsertEdge(newEdgeVal, nodeFromVal, nodeToVal interface{}) {
	nodes := map[interface{}]*Node{nodeFromVal: nil, nodeToVal: nil}
	for _, node := range r.Nodes {
		if _, ok := nodes[node.Value]; ok {
			nodes[node.Value] = node
		}
	}

}
