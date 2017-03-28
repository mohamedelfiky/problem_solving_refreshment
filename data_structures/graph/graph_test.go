package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	g := NewGraph()

	assert.Equal(t, -1, g.findMaxIndex())

	g.InsertEdge(51, 0, 1) // MV <-> SF

	assert.Equal(t, 1, g.findMaxIndex())

	g.SetNodeNames([]string{"Mountain View", "San Francisco", "London", "Shanghai", "Berlin", "Sao Paolo", "Bangalore"})

	g.InsertEdge(51, 1, 0)    // SF <-> MV
	g.InsertEdge(9950, 0, 3)  // MV <-> Shanghai
	g.InsertEdge(9950, 3, 0)  // Shanghai <-> MV
	g.InsertEdge(10375, 0, 5) // MV <-> Sao Paolo
	g.InsertEdge(10375, 5, 0) // Sao Paolo <-> MV
	g.InsertEdge(9900, 1, 3)  // SF <-> Shanghai
	g.InsertEdge(9900, 3, 1)  // Shanghai <-> SF
	g.InsertEdge(9130, 1, 4)  // SF <-> Berlin
	g.InsertEdge(9130, 4, 1)  // Berlin <-> SF
	g.InsertEdge(9217, 2, 3)  // London <-> Shanghai
	g.InsertEdge(9217, 3, 2)  // Shanghai <-> London
	g.InsertEdge(932, 2, 4)   // London <-> Berlin
	g.InsertEdge(932, 4, 2)   // Berlin <-> London
	g.InsertEdge(9471, 2, 5)  // London <-> Sao Paolo
	g.InsertEdge(9471, 5, 2)  // Sao Paolo <-> London

	adjList, adjListNames := g.GetAdjacencyList(), g.GetAdjacencyListNamesOnly()

	assert.Equal(t, g.Nodes[0].Edges[0].NodeTo, adjList[0][0])
	assert.Equal(t, adjListNames["Mountain View"], []string{"San Francisco", "Shanghai", "Sao Paolo"})
	assert.Equal(t, adjListNames["San Francisco"], []string{"Mountain View", "Shanghai", "Berlin"})
	assert.Equal(t, adjListNames["Sao Paolo"], []string{"Mountain View", "London"})
	assert.Equal(t, adjListNames["Shanghai"], []string{"Mountain View", "San Francisco", "London"})
	assert.Equal(t, adjListNames["London"], []string{"Shanghai", "Berlin", "Sao Paolo"})
	adjMatrix := g.GetAdjacencyMatrix()
	assert.Equal(t, 51, adjMatrix[0][1])
	assert.Equal(t, 9471, adjMatrix[5][2])
	assert.Equal(t, 9900, adjMatrix[1][3])

	assert.Equal(t, []string{"London", "Shanghai", "Mountain View", "San Francisco", "Berlin", "Sao Paolo"}, g.DfsNamesOnly(2))
	assert.Equal(t, []string{"London", "Shanghai", "Berlin", "Sao Paolo", "Mountain View", "San Francisco"}, g.BFSNamesOnly(2))
}
