package main

import "fmt"

/**
Insert sort 适用场景：
 1）数组中重复数据占比比较大时
 2）数组近乎有序时
*/

func InsertSort(arr []int) {
	leng := len(arr)
	for i := 1; i < leng; i++ {
		for j := i - 1; j >= 0 && (arr[j+1] < arr[j]); j-- {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func InsertSort2(arr []int) {

	leng := len(arr)
	for i := 1; i < leng; i++ {

		tmp := arr[i]
		j := i
		for ; j > 0 && (arr[j-1] > tmp); j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}

}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	InsertSort(arr)
	fmt.Println(arr)
	arr2 := []int{5, 4, 3, 2, 1}
	InsertSort2(arr2)
	fmt.Println(arr2)
}
