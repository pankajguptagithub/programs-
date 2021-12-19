package main

import (
	"fmt"
)
func GenDisplaceFn(a float64,v float64,s0 float64) func(float64)float64 {
	fn := func(t float64)float64{
			return ((0.5*a*t*t)+(v*t)+(s0))
	}
	return fn
}
func main(){
	var a,v,s0,t float64
	fmt.Println("Enter acceleration ")
	fmt.Scanf("%f",&a)
	fmt.Println("Enter initial velocity")
	fmt.Scanf("%f",&v)
	fmt.Println("Enter initial distance travelled")
	fmt.Scanf("%f",&s0)
	fmt.Println("Enter time")
	fmt.Scanf("%f",&t)
	var fn func(float64)float64
	fn = GenDisplaceFn(a,v,s0)
	fmt.Println(fn(t))
}