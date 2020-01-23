package adjacency_list

import (
	list "container/list"
	"fmt"
)

type AdjacencyList (map[interface{}]*list.List)

func (al *AdjacencyList) AddVertex(v1 interface{}) {
	if _, found := (*al)[v1]; !found {
		(*al)[v1] = list.New()
	}
}

func (al *AdjacencyList) AddEdge(v1 interface{}, v2 interface{}) {
	if _, found := (*al)[v1]; !found {
		al.AddVertex(v1)
	}
	((*al)[v1]).PushFront(v2)

	if _, found := (*al)[v2]; !found {
		al.AddVertex(v2)
	}
	((*al)[v2]).PushFront(v1)
}

func (al *AdjacencyList) Print() {
	fmt.Println("AdjacencyList:")
	for k, v := range *al {
		fmt.Printf("vertex -> [%+v] :", k)

		ele := (*v).Front()
		for ; ele != nil; ele = ele.Next() {
			fmt.Printf(" -> %v", (*ele).Value)
		}
		fmt.Println()
	}
}
