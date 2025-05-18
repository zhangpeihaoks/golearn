package quick_sort

import (
	"math/rand/v2"
)

func QuickSort(arr []int, asc bool) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]

	var left, right, mid []int

	for _, v := range arr {
		if asc {
			switch {
			case v < pivot:
				left = append(left, v)
			case v == pivot:
				mid = append(mid, v)
			case v > pivot:
				right = append(right, v)
			}
		} else {
			switch {
			case v < pivot:
				right = append(right, v)
			case v == pivot:
				mid = append(mid, v)
			case v > pivot:
				left = append(left, v)
			}
		}

	}

	right = QuickSort(right, asc)
	left = QuickSort(left, asc)

	arr = append(left, mid...)
	arr = append(arr, right...)

	return arr
}

func QuickSort2(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := arr[selectPivot(left, right)]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	QuickSort2(arr, left, i)
	QuickSort2(arr, i+2, right)
}

func selectPivot(left, right int) int {
	return rand.IntN(right-left) + left
}
