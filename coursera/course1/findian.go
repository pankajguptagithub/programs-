package main

import (
	"fmt"i
	"strings"
)

func main(){
	var str string
	fmt.Println("Enter any string")
	fmt.Scan(&str)
    if strings.Contains(str,"i") && strings.Contains(str,"a") && strings.Contains(str,"n"){
		fmt.Println("Found")
	}else{
		fmt.Println("Not found")
	}
}