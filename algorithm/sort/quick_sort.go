package sort

import (
	"context"
	"math/rand/v2"
)

func QuickSort(ctx context.Context, arr []int, asc bool) {
	quickSort(ctx, arr, asc, 0, len(arr)-1)
}

func quickSort(ctx context.Context, arr []int, asc bool, left, right int) {
	if left >= right {
		return
	}
	select {
	case <-ctx.Done():
		return
	default:
	}
	pivot := arr[selectPivot(left, right)]
	i, j := left, right-1

	compare := func(a, b int) bool {
		if asc {
			return a <= b
		}
		return a >= b
	}

	for i <= j {
		for i <= j && compare(arr[i], pivot) {
			i++
		}
		for i <= j && !compare(arr[j], pivot) {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}

	}

	arr[i], arr[right] = arr[right], arr[i]

	quickSort(ctx, arr, asc, left, i-1)
	quickSort(ctx, arr, asc, i+1, right)
}

func QuickSort2(ctx context.Context, arr []int, asc bool) {
	type Range struct {
		Left, Right int
	}

	stack := []Range{{0, len(arr) - 1}}

	compare := func(a, b int) bool {
		if asc {
			return a <= b
		}
		return a >= b
	}

	for len(stack) > 0 {
		select {
		case <-ctx.Done():
			return
		default:

		}

		n := len(stack) - 1
		left, right := stack[n].Left, stack[n].Right
		stack = stack[:n]

		if left >= right {
			continue
		}

		pivot := arr[selectPivot(left, right)]
		i, j := left, right-1

		for i <= j {

			for i <= j && compare(arr[i], pivot) {
				i++
			}
			for i <= j && !compare(arr[j], pivot) {
				j--
			}

			if i < j {
				arr[i], arr[j] = arr[j], arr[i]
			}

		}

		arr[i], arr[right] = arr[right], arr[i]

		stack = append(stack, Range{left, i - 1}, Range{i + 1, right})
	}
}

func QuickSort3(ctx context.Context, arr []int, asc bool) {
	copy(arr, quickSort3(ctx, arr, asc))
}

func quickSort3(ctx context.Context, arr []int, asc bool) []int {
	if len(arr) < 2 {
		return arr
	}
	select {
	case <-ctx.Done():
		return arr
	default:
	}
	pivot := arr[selectPivot(0, len(arr)-1)]
	var left, mid, right []int
	for _, num := range arr {

		if asc {
			switch {
			case num < pivot:
				left = append(left, num)
			case num == pivot:
				mid = append(mid, num)
			case num > pivot:
				right = append(right, num)
			}
		} else {
			switch {
			case num > pivot:
				left = append(left, num)
			case num == pivot:
				mid = append(mid, num)
			case num < pivot:
				right = append(right, num)
			}
		}
	}
	left = quickSort3(ctx, left, asc)
	right = quickSort3(ctx, right, asc)
	return append(left, append(mid, right...)...)
}

func selectPivot(left, right int) int {
	return rand.IntN(right-left) + left
}
