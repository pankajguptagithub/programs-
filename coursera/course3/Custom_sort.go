package main

import (
	"fmt"
	"sync"
)

func Swap(arr []int, i int) {
	var temp int
	temp = arr[i+1]
	arr[i+1] = arr[i]
	arr[i] = temp
}
func BubbleSort(wg *sync.WaitGroup, arr []int) {
	defer wg.Done()
	for i := 0; i < len(arr); i++ {
		for j := 0; j+i < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
	fmt.Println(arr)
}
func combine_arrays(a []int, b []int) []int {
	fmt.Println("1st array to combine is ", a)
	fmt.Println("2nd array to combine is ", b)
	final_array := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final_array = append(final_array, a[i])
			i++
		} else {
			final_array = append(final_array, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final_array = append(final_array, a[i])
	}
	for ; j < len(b); j++ {
		final_array = append(final_array, b[j])
	}

	fmt.Println("Array after combination is ", final_array)
	return final_array
}
func merge_arrays(arr1 []int, arr2 []int, arr3 []int, arr4 []int) []int {
	result1 := combine_arrays(arr1, arr2)
	fmt.Println(result1)
	result2 := combine_arrays(result1, arr3)
	fmt.Println(result1)
	result3 := combine_arrays(result2, arr4)
	fmt.Println(result1)
	return result3
}
func main() {
	var wg sync.WaitGroup
	fmt.Println("Enter the number of elements in the array")
	var n int
	fmt.Scan(&n)
	var arr = make([]int, n)
	fmt.Println("Enter the array elements")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("Array entered by you", arr)
	batches := 4
	batch_size := (len(arr) + batches - 1) / batches
	arr1 := arr[0:batch_size]
	arr2 := arr[batch_size*1 : batch_size*2]
	arr3 := arr[batch_size*2 : batch_size*3]
	arr4 := arr[batch_size*3 : n]

	wg.Add(4)
	go BubbleSort(&wg, arr1)
	go BubbleSort(&wg, arr2)
	go BubbleSort(&wg, arr3)
	go BubbleSort(&wg, arr4)
	wg.Wait()
	final_array := merge_arrays(arr1, arr2, arr3, arr4)
	fmt.Println("Array elements in sorted order ", final_array)
}
