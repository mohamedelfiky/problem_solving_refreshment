// implementation of BFS and DFS graph traversing methods
package graph

import "github.com/mohamedelfiky/problem_solving_refreshment/data_structures/queue"

type Node struct {
	Value   int
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
	nodeMap   map[int]*Node
}

func NewGraph() *Graph {
	return &Graph{Nodes: []*Node{}, Edges: []*Edge{}, NodeNames: []string{}, nodeMap: map[int]*Node{}}
}

func (r *Graph) SetNodeNames(nodeNames []string) {
	r.NodeNames = nodeNames
}

func (r *Graph) InsertNode(newNodeVal int) *Node {
	node := Node{Value: newNodeVal}
	r.Nodes = append(r.Nodes, &node)
	r.nodeMap[node.Value] = &node
	return &node
}

func (r *Graph) InsertEdge(newEdgeVal interface{}, nodeFromVal, nodeToVal int) {
	nodes := map[int]*Node{nodeFromVal: nil, nodeToVal: nil}
	for _, node := range r.Nodes {
		if _, ok := nodes[node.Value]; ok {
			nodes[node.Value] = node
		}
		if v1, v2 := nodes[nodeFromVal], nodes[nodeToVal]; v1 != nil && v2 != nil {
			break
		}
	}

	for val, node := range nodes {
		if node == nil {
			nodes[val] = r.InsertNode(val)
		}
	}

	nodeFrom, nodeTo := nodes[nodeFromVal], nodes[nodeToVal]
	newEdge := Edge{Value: newEdgeVal, NodeFrom: nodeFrom, NodeTo: nodeTo}

	nodeFrom.Edges = append(nodeFrom.Edges, &newEdge)
	nodeTo.Edges = append(nodeTo.Edges, &newEdge)
	r.Edges = append(r.Edges, &newEdge)
}

func (r *Graph) GetAdjacencyList() map[int][]*Node {
	res := map[int][]*Node{}
	for _, edge := range r.Edges {
		if _, ok := res[edge.NodeFrom.Value]; ok {
			res[edge.NodeFrom.Value] = append(res[edge.NodeFrom.Value], edge.NodeTo)
		} else {
			res[edge.NodeFrom.Value] = []*Node{edge.NodeTo}
		}
	}
	return res
}

func (r *Graph) GetAdjacencyListNamesOnly() map[string][]string {
	res := map[string][]string{}
	for _, edge := range r.Edges {
		node_name := r.NodeNames[edge.NodeFrom.Value]
		if _, ok := res[node_name]; ok {
			res[node_name] = append(res[node_name], r.NodeNames[edge.NodeTo.Value])
		} else {
			res[node_name] = []string{r.NodeNames[edge.NodeTo.Value]}
		}
	}
	return res
}

func (r *Graph) GetAdjacencyMatrix() [][]interface{} {
	matrix := make([][]interface{}, r.findMaxIndex())
	for i := 0; i < r.findMaxIndex(); i++ {
		matrix[i] = make([]interface{}, r.findMaxIndex())
	}

	for _, edge := range r.Edges {
		matrix[edge.NodeFrom.Value][edge.NodeTo.Value] = edge.Value
	}
	return matrix
}

func (r *Graph) DFS(startNode int) []int {
	r.clearVisited()
	return dfsHelper(r.FindNode(startNode))
}

func (r *Graph) BFS(startNodeValue int) []int {
	r.clearVisited()
	startNode, res := r.FindNode(startNodeValue), []int{startNodeValue}
	startNode.Visited = true
	q := queue.Queue{}

	for len(res) < len(r.Nodes) {
		for _, edge := range startNode.Edges {
			if !edge.NodeTo.Visited {
				edge.NodeTo.Visited = true
				res = append(res, edge.NodeTo.Value)
				q.Push(edge.NodeTo)
			}
		}
		startNode = q.Pop().(*Node)
	}

	return res
}

func (r *Graph) BFSNamesOnly(startNode int) []string {
	return r.getNodeNames(r.BFS(startNode))
}

func (r *Graph) DfsNamesOnly(startNode int) []string {
	return r.getNodeNames(r.DFS(startNode))
}

func (r *Graph) FindNode(value int) *Node {
	return r.nodeMap[value]
}

func (r *Graph) getNodeNames(nodeValues []int) []string {
	res := make([]string, len(nodeValues))
	for i, node := range nodeValues {
		res[i] = r.NodeNames[node]
	}
	return res
}

func dfsHelper(startNode *Node) []int {
	nodes := []int{startNode.Value}
	startNode.Visited = true
	for _, edge := range startNode.Edges {
		if !edge.NodeTo.Visited {
			for _, n := range dfsHelper(edge.NodeTo) {
				nodes = append(nodes, n)
			}
		}
	}
	return nodes
}

func (r *Graph) clearVisited() {
	for _, node := range r.Nodes {
		node.Visited = false
	}
}

func (r *Graph) findMaxIndex() int {
	if len(r.NodeNames) > 0 {
		return len(r.NodeNames)
	}
	max_index := -1
	if len(r.Nodes) > 0 {
		for _, node := range r.Nodes {
			if node.Value > max_index {
				max_index = node.Value
			}
		}
	}
	return max_index
}
