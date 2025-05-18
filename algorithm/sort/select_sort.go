package sort

import "context"

func SelectSort(ctx context.Context, arr []int, asc bool) {
	for i := 0; i < len(arr); i++ {
		select {
		case <-ctx.Done():
			return
		default:
		}
		targetIndex := i
		for j := i + 1; j < len(arr); j++ {

			if (asc && arr[j] <= arr[targetIndex]) || (!asc && arr[j] > arr[targetIndex]) {
				targetIndex = j
			}

		}
		if targetIndex != i {
			arr[i], arr[targetIndex] = arr[targetIndex], arr[i]
		}
	}
}
