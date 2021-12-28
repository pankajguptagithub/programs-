package demoapi

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"os"

	"github.com/gorilla/mux"
)
type details struct {
    name      string    `json:"name"`
    address   string     `json:"address"`
    city      string `json:"city"`
}
func add_details(response http.ResponseWriter, request *http.Request) {
	type all_details []details
	vars := mux.Vars(request)
	input_name := vars["name"]
	input_address := vars["address"]
	input_cty := vars["city"]
	var all_details = []details{details{name: input_name , address: input_address, city: input_city}}
	w.Write(o[]byte("Details entered by you - "+all_details))
	f, err := os.Open("savedetails.json")
	if err != nil {
        	panic(err)
    	}
	defer func() {
        if err := f.Close(); err != nil {
            	panic(err)
        	}
    	}()
	f.Write(all_details)
}	
