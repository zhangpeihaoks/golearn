package greedy

import (
	"sort"
)

// OptimalShipment w 每个集装箱的重量 x 货物的装载情况 c 集装箱的最大承载量
func OptimalShipment(w []int, c int) []int {
	sort.Ints(w)
	selectedW := make([]int, 0)
	for len(w) > 0 && c > 0 {
		currentW := w[0]
		if currentW <= c {
			selectedW = append(selectedW, currentW)
			c -= currentW
			w = w[1:]
		} else {
			break
		}
	}
	return selectedW
}
