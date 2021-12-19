package main 

import ("fmt")

func main(){

	var i float64
	fmt.Println("Enter a number")
	fmt.Scan(&i)
	fmt.Println("Entered number is",i)
	fmt.Println("Integer converted number is",int32(i))
}