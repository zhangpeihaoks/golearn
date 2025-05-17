package graph

import "log/slog"

func Dfs(graph []*IGraphNode, depth int, visited map[*IGraphNode]bool, order *[]int) {
	for _, node := range graph {
		if visit, ok := visited[node]; ok && visit {
			continue
		}
		visited[node] = true
		slog.Info("node", "num", node.GetNum(), "depth", depth)
		*order = append(*order, node.GetNum())
		if len(node.NodeMap) > 0 {
			list := func() []*IGraphNode {
				list := make([]*IGraphNode, 0)
				for k := range node.NodeMap {
					list = append(list, k)
				}
				return list
			}
			Dfs(list(), depth+1, visited, order)
		}
	}
}
