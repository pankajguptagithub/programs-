package main

import (
	"fmt"
	"net/http"
	//"encoding/json"
	"os"

	"github.com/gorilla/mux"
)
func err_check(response http.ResponseWriter,err error){
	 if err != nil {
                fmt.Fprintf(response,err.Error())
	 }	
}
func Collect(response http.ResponseWriter, request *http.Request) {
	//vars := mux.Vars(request)
	//playerName := vars["a"]
	//fmt.Println(d)
	//jsondata,err := json.Marshal(d)
	//err_check(response,err)
	//fmt.Println(string(jsondata))
	//err_check(response,err)
	file,_ := os.ReadFile("collectdata.json")
	fmt.Fprintf(response,string(file))
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/collect", Collect).Methods("GET")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err)
	}
}
