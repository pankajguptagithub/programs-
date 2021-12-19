package main

import (
	"fmt"
)
func Swap(arr []int,i int){
	var temp int
	temp = arr[i+1]
	arr[i+1] = arr[i]
	arr[i] = temp
}
func BubbleSort(arr []int){
	fmt.Println("Size of the array is ",len(arr))
	for i:=0;i<len(arr);i++{
		for j:=0;j+i<len(arr)-1;j++{
			if (arr[j]>arr[j+1]){
				Swap(arr,j)
			}
		}
	}
}
func main(){
	fmt.Println("Enter the number of elements in the array")
	var n int
	fmt.Scan(&n)
	var arr = make([]int,n)	
	fmt.Println("Enter the array elements") 
	for i:=0;i<n;i++{
		fmt.Scan(&arr[i])
	}
	fmt.Println("Array entered by you",arr)
    BubbleSort(arr)
	fmt.Println("Array elements in sorted order ",arr)
}