package main

import (
    "os"
    "net/http"
    "io"
    "fmt"
    "encoding/json"
)
type image_download struct{
	name string
	size int
}
func error_check(err error){
if err != nil  {
	fmt.Println("Error encountered - ",err)
	panic(err)
  }
}
func download_file(w http.ResponseWriter,r *http.Request){
	image_url := r.URL.Query()["image"]
	out, err := os.Create("pic.jpg")
	error_check(err)
	defer out.Close()

	resp, err := http.Get(image_url[0])
	error_check(err)
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	error_check(err)

	fileStat, err := os.Stat("pic.jpg")
	error_check(err)

	id := image_download{name: fileStat.Name(),size: int(fileStat.Size()),}
	output,_  := json.Marshal(id)
	op := string(output)

	var id image_download
	id.name = fileStat.Name()
	id.size =  int(fileStat.Size())

	output,err := json.Marshal(id)
	error_check(err)

	w.Write([]byte(op))
}
func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		download_file(w,r)
	})
	http.ListenAndServe(":3000", nil)
}
