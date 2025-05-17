package graph

type IGraphNode struct {
	Num     int
	NodeMap map[*IGraphNode]int
}

func (g *IGraphNode) GetNum() int {
	return g.Num
}

func (g *IGraphNode) GetNodeList() map[*IGraphNode]int {
	return g.NodeMap
}

func (g *IGraphNode) AddNode(node *IGraphNode, weight int) {
	g.NodeMap[node] = weight
}

func (g *IGraphNode) RemoveNode(node *IGraphNode) {
	delete(g.NodeMap, node)
}

func NewNode(num int) *IGraphNode {
	return &IGraphNode{
		Num:     num,
		NodeMap: make(map[*IGraphNode]int),
	}
}

func NewGraph(nums []int, edges [][]int) []*IGraphNode {
	graph := make([]*IGraphNode, len(nums))
	for index, num := range nums {
		graph[index] = NewNode(num)
	}
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		graph[from].AddNode(graph[to], weight)
	}
	return graph

}
