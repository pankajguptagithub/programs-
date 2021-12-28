package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"os"

	"github.com/gorilla/mux"
)
type details struct {
    Name string		 `json: "name"`
    City string 	 `json: "city"`
}
func err_check(response http.ResponseWriter,err error){
	 if err != nil {
                fmt.Fprintf(response,err.Error())
	 }	
}
func Collect(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	d := details{Name: vars["a"],City: vars["b"]}
	fmt.Fprintf(response,"writing data to file collectdata.json")
	fmt.Println(d)
	jsondata,err := json.Marshal(d)
	err_check(response,err)
	fmt.Println(string(jsondata))
	err_check(response,err)
	file, err := os.OpenFile("collectdata.json", os.O_APPEND|os.O_WRONLY, 0644)
	err_check(response,err)
	defer file.Close()
	_,err = file.WriteString(string(jsondata))
	err_check(response,err)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/collect/{a}/{b}", Collect).Methods("GET")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err)
	}
}
