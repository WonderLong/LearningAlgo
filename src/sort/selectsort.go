package main

import "fmt"

type Sortable interface {
	Length() int
	LessThan(int, int) bool
	Swap(int, int)
}

func SelectSort(arr Sortable) {
	leng := arr.Length()
	for i := 0; i < leng; i++ {
		minIdx := i
		for j := i + 1; j < leng; j++ {
			if arr.LessThan(j, minIdx) {
				minIdx = j
			}
		}
		arr.Swap(i, minIdx)
	}
}

type IntArray []int

func (arr IntArray) Length() int {
	return len(arr)
}

func (arr IntArray) LessThan(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr IntArray) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	arr := IntArray{5, 4, 3, 2, 1}
	SelectSort(arr)
	fmt.Println(arr)
}
