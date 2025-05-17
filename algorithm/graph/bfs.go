package graph

import "log/slog"

func Bfs(graph []*IGraphNode, visited map[*IGraphNode]bool, order *[]int) {
	queue := make([]*IGraphNode, 0)

	if len(graph) > 0 {
		queue = append(queue, graph[0])
		visited[graph[0]] = true
	}

	for len(queue) > 0 {
		current := queue[0]

		slog.Info("node", "num", current.GetNum())
		*order = append(*order, current.GetNum())
		for neighbor := range current.NodeMap {
			if _, ok := visited[neighbor]; !ok {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
		queue = queue[1:] // Dequeue the first element
	}
}
