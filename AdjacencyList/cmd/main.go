package main

import (
	al "github.com/pramineni01/graph/AdjacencyList/adjacency_list"
)

func main() {
	al := make(al.AdjacencyList, 0)

	// add vertices 1,2,3
	al.AddVertex(1)
	al.AddVertex(2)
	al.AddVertex(3)

	al.AddEdge(1, 2)
	al.AddEdge(2, 3)
	//al.AddEdge(3, 1)

	al.Print()
}
