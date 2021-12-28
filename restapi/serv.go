package main

import (
    "net/http"
)
type HttpHandler struct{}
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    data := []byte("Hello Pankaj!") 
    res.Write(data)
}

func main() {
    handler := HttpHandler{}
    http.ListenAndServe(":9000", handler)
}
