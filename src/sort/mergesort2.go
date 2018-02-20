package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 4, 3, 2, 1}
	MergeSort(arr)
	fmt.Println(arr)
}

//自底向上的归并排序
func MergeSort(arr []int) {
	doMergeSort(arr)
}

func doMergeSort(arr []int) {
	n := len(arr)
	for sz := 1; sz <= n; sz += sz {

		for i := 0; i+sz < n; i += sz + sz {

			if arr[i+sz-1] > arr[i+sz] {
				fmt.Printf("sz: %d; i: %d \r\n", sz, i)
				merge(arr, i, i+sz-1, min(i+sz+sz-1, n-1))
			}

		}

	}

}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
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
