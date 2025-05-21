package sort

func TestQuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	pivot := arr[right]
	i, j := left, right-1
	for i <= j {
		for i <= j && arr[i] < pivot {
			i++
		}
		for i <= j && arr[j] > pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i], arr[right] = arr[right], arr[i]

	TestQuickSort(arr, left, i-1)
	TestQuickSort(arr, i+1, right)

}
