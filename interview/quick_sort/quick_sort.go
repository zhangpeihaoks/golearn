package main

import "log/slog"

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var left, mid, right []int
	for i := 0; i < len(arr); i++ {
		switch {
		case arr[i] < pivot:
			left = append(left, arr[i])
		case arr[i] == pivot:
			mid = append(mid, arr[i])
		case arr[i] > pivot:
			right = append(right, arr[i])
		}
	}
	right = QuickSort(right)
	left = QuickSort(left)
	arr = append(left, mid...)
	arr = append(arr, right...)
	return arr
}

func main() {
	var arr = []int{2, 3, 4, 5, 65, 11, 22, 1}
	arr = QuickSort(arr)
	slog.Info("data", "arr", arr)
}
