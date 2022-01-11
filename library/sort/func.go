package sort

// SelectSort 选择排序
func SelectSort(arr []int)  {
	for i := 0; i < len(arr); i++ {
		var pos = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[pos] {
				pos = j
			}
		}
		arr[i], arr[pos] = arr[pos], arr[i]
	}
}


// BubblingSort 冒泡排序

func BubblingSort(arr []int)  {
	for i := len(arr) - 1; i > 0 ; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
}
