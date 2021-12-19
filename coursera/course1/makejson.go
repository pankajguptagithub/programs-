package main

import (
	"fmt"
	"encoding/json"
)
func main(){
	user := map[string]string{}
	var name string
	var address string
	fmt.Println("Enter user name")
	fmt.Scan(&name)
	fmt.Println("Enter user address")
	fmt.Scan(&address)
	user["name"] = name
	user["address"] = address
	out,_ := json.Marshal(user)
	fmt.Println(string(out))
}