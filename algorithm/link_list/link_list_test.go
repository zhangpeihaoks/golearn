package link_list

import (
	"fmt"
	"golearn/algorithm/quick_sort"
	"math/rand"
	"testing"
	"time"
)

func TestLinkList(t *testing.T) {
	inputData := make([]int, 100)
	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 100; i++ {
		inputData[i] = rand.Intn(100)
	}
	inputData = quick_sort.QuickSort(inputData, true)

	headNode := NewNode[int](nil, inputData[0])
	tailNode := headNode
	for i := 1; i < len(inputData); i++ {
		tmpNode := NewNode(tailNode, inputData[i])
		tailNode = tmpNode
	}

	nowNode := headNode
	index := 0
	for {
		fmt.Println("index", index, "nowNode.Next:", nowNode.Next, "data:", nowNode.Value)
		index++
		if nowNode.Next == nil {
			break
		}
		nowNode = nowNode.Next
	}

}
