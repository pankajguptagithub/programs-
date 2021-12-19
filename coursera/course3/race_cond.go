package main

import ("fmt")

func incr(y *int){
	*y = *y + 4
}
func incrr(z *int){
	*z = *z + 5
}
func main(){
	x := 1
	go incr(&x)
	go incrr(&x)
	fmt.Println("Final value of x is ",x)
}
