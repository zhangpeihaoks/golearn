package sort

import "context"

func ShellSort(ctx context.Context, arr []int, asc bool) {
	n := len(arr)
	gap := n / 2
	for gap > 0 {
		for i := gap; i < n; i++ {
			tmp := arr[i]
			j := i
			select {
			case <-ctx.Done():
				return
			default:
			}
			for j >= gap && ((asc && arr[j-gap] > tmp) || (!asc && arr[j-gap] < tmp)) {

				arr[j] = arr[j-gap]
				j -= gap

			}
			arr[j] = tmp
		}
		gap /= 2
	}
}
