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

// InsertSort 插入排序
func InsertSort(arr []int)  {
	for i := 1;i < len(arr);i++ {
		for j := i; j > 0; j-- {
			if arr[j] > arr[j - 1] {
				arr[j], arr[j - 1] = arr[j - 1], arr[j]
			}
		}
	}
}

// ShellSort 希尔排序
func ShellSort(arr []int)  {
	var h = 1
	for h <= len(arr)/3 {
		h = h*3 + 1
	}
	for gap := h;gap > 0;gap = (gap - 1)/3 {
		for i := gap;i < len(arr);i++ {
			for j := i; j > gap - 1; j -= gap {
				if arr[j] < arr[j - gap] {
					arr[j], arr[j - gap] = arr[j - gap], arr[j]
				}
			}
		}
	}
}

func MergeSort(arr []int, left, right int)  {
	if left == right {
		return
	}
	//分成两部分，左右排序，左右merge
	var mid = left + (right-left)/2

	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)

	merge(arr, left, mid+1, right)
}

func merge(arr []int, leftPtr, rightPtr, rightBound int) {
	var mid = rightPtr - 1
	var newArr = make([]int, rightBound-leftPtr+1)

	var i, j, k = leftPtr, rightPtr, 0

	for i <= mid && j <= rightBound {
		if arr[i] <= arr[j] {
			newArr[k] = arr[i]
			k++
			i++
		} else {
			newArr[k] = arr[j]
			k++
			j++
		}
	}

	for i <= mid {
		newArr[k] = arr[i]
		k++
		i++
	}

	for j <= rightBound {
		newArr[k] = arr[j]
		k++
		j++
	}
	for ii := 0; ii < len(newArr); ii++ {
		arr[leftPtr+ii] = newArr[ii]
	}
}

func QuickSort(arr []int, left , right int)  {
	if left >= right {
		return
	}
	mid := partition(arr, left, right)
	QuickSort(arr, left, mid-1)
	QuickSort(arr, mid+1, right)
}

func partition(arr []int, leftBound, rightBound int) int {
	var pivot = arr[rightBound]
	var left, right = leftBound, rightBound - 1
	for left <= right {
		for left <= right && arr[left] <= pivot {
			left++
		}
		for left <= right && arr[right] > pivot {
			right--
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[left], arr[rightBound] = arr[rightBound], arr[left]
	return left
}

func CountSort(arr []int) []int  {
	var count = make([]int, 10)
	var result = make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	for i := 1; i < len(count); i++ {
		count[i] = count[i] + count[i-1]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		count[arr[i]] -= 1
		result[count[arr[i]]] = arr[i]
	}

	return result
}