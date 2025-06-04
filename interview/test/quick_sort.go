package test

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	// 选择最左边的元素作为基准值
	pivot := arr[left]
	i, j := left, right
	for i <= j {
		for i <= j && arr[i] <= pivot {
			i++
		}
		for i <= j && arr[j] >= pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// 将基准值放到正确的位置
	arr[left], arr[j] = arr[j], arr[left]
	quickSort(arr, left, j-1)
	quickSort(arr, j+1, right)
}
