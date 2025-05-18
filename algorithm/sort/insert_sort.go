package sort

import "context"

func InsertSort(ctx context.Context, list []int, asc bool) {
	for i := 1; i < len(list); i++ {
		select {
		case <-ctx.Done():
			return
		default:
		}
		tmp := list[i]
		j := i - 1
		for j >= 0 && ((asc && list[j] > tmp) || (!asc && list[j] <= tmp)) {

			list[j+1] = list[j]
			j--

		}
		list[j+1] = tmp
	}
}
