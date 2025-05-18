package sort

import "context"

func BubbleSort(ctx context.Context, arr []int, asc bool) {
	for i := 0; i < len(arr); i++ {
		select {
		case <-ctx.Done():
			return
		default:
		}
		for j := 0; j < len(arr)-i-1; j++ {

			if (asc && arr[j] > arr[j+1]) || (!asc && arr[j] < arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}

		}
	}

}

func BubbleSort2(ctx context.Context, arr []int, asc bool) {
	isSorted := true
	for i := 0; i < len(arr) && isSorted; i++ {
		select {
		case <-ctx.Done():
			return
		default:
		}
		isSorted = false
		for j := 0; j < len(arr)-i-1; j++ {

			if (asc && arr[j] > arr[j+1]) || (!asc && arr[j] < arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				isSorted = true
			}

		}
	}

}
