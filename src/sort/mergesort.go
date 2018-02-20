package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	MergeSort(arr)
	fmt.Println(arr)
}

/***

**/
func MergeSort(arr []int) {
	doMergeSort(arr, 0, len(arr)-1)
}

func doMergeSort(arr []int, s int, e int) {

	if s >= e {
		return
	}

	//一般的mid = (s + e)/2, 为避免中间值越界的写法：
	mid := s + (e-s)/2

	doMergeSort(arr, s, mid)
	doMergeSort(arr, mid+1, e)
	if arr[mid] > arr[mid+1] {
		merge(arr, s, mid, e)
	}
}

//merge arr[s..mid] and arr[mid+1..e]两个都是闭区间
func merge(arr []int, s, mid, e int) {
	//要注意合并过程的边界
	auxLength := e - s + 1
	aux := make([]int, auxLength)
	for i := s; i <= e; i++ {
		aux[i-s] = arr[i]
	}

	i := s
	j := mid + 1
	for k := s; k <= e; k++ {

		if i > mid {
			arr[k] = aux[j-s]
			j = j + 1
		} else if j > e {
			arr[k] = aux[i-s]
			i = i + 1
		} else if aux[i-s] < aux[j-s] {
			arr[k] = aux[i-s]
			i = i + 1
		} else {
			arr[k] = aux[j-s]
			j = j + 1
		}
	}
}
