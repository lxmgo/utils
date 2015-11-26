// 快速排序
//
// Copyright (c) 2015 - Batu <1235355@qq.com>
//
package quick

//快速排序方式一
func quickSort(values []int, left int, right int) {
	key := values[left] //取出第一项
	p := left
	i, j := left, right

	for i <= j {
		//由后开始向前搜索(j--)，找到第一个小于key的值values[j]
		for j >= p && values[j] >= key {
			j--
		}

		//第一个小于key的值 赋给 values[p]
		if j >= p {
			values[p] = values[j]
			p = j
		}

		if values[i] <= key && i <= p {
			i++
		}

		if i < p {
			values[p] = values[i]
			p = i
		}

		values[p] = key
		if p-left > 1 {
			quickSort(values, left, p-1)
		}
		if right-p > 1 {
			quickSort(values, p+1, right)
		}
	}

}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}

//快速排序方式二
func quickSort2(sortArray []int, left, right int) {
	if left < right {
		key := sortArray[(left+right)/2]
		i := left
		j := right

		for {
			for sortArray[i] < key {
				i++
			}
			for sortArray[j] > key {
				j--
			}
			if i >= j {
				break
			}
			sortArray[i], sortArray[j] = sortArray[j], sortArray[i]
		}

		quickSort2(sortArray, left, i-1)
		quickSort2(sortArray, j+1, right)
	}
}

func QuickSort2(values []int) {
	quickSort2(values, 0, len(values)-1)
}
