package quick_sort

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
