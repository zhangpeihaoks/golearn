package graph

import (
	"fmt"
	"log/slog"
	"testing"
)

func TestNewGraph(t *testing.T) {
	nums := []int{
		0, 1, 2, 3, 4, 5,
	}
	edges := [][]int{
		{0, 1, 1},
		{0, 2, 1},
		{1, 3, 1},
		{2, 5, 1},
		{1, 4, 1},
	}
	graph := NewGraph(nums, edges)
	for i, node := range graph {
		slog.Info("node", "index", i, "num", node.GetNum(), " node", node.NodeMap)
	}

	fmt.Println()
	var dfsOrder []int
	Dfs(graph, 0, make(map[*IGraphNode]bool), &dfsOrder)
	fmt.Println("DFS order:", dfsOrder)

	fmt.Println()
	var bfsOrder []int
	Bfs(graph, make(map[*IGraphNode]bool), &bfsOrder)
	fmt.Println("BFS order:", bfsOrder)
}
