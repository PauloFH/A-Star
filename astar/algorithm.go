package astar

import (
	"container/heap"

	"github.com/PauloFH/A-Star/data"
)

func FindPath(start, goal string) ([]string, int) {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	hStart, ok := data.Heuristics[start]
	if !ok {
		hStart = 0
	}

	heap.Push(&pq, &Item{City: start, Priority: 0 + hStart, G: 0})
	visited := make(map[string]bool)

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item)

		if current.City == goal {
			return reconstructPath(current)
		}

		if visited[current.City] {
			continue
		}
		visited[current.City] = true
		if edges, ok := data.Graph[current.City]; ok {
			for _, edge := range edges {
				if !visited[edge.To] {
					newG := current.G + edge.Cost

					hNext, ok := data.Heuristics[edge.To]
					if !ok {
						hNext = 0
					}

					newF := newG + hNext

					heap.Push(&pq, &Item{
						City:     edge.To,
						Priority: newF,
						G:        newG,
						Parent:   current,
					})
				}
			}
		}
	}

	return []string{}, 0
}

func reconstructPath(node *Item) ([]string, int) {
	cost := node.G
	var path []string
	for node != nil {
		path = append([]string{node.City}, path...)
		node = node.Parent
	}
	return path, cost
}
