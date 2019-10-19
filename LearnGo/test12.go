//HTTP Server
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		w.Write([]byte("Received a GET request.\n"))
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, v)
		}
	case "POST":
		w.Write([]byte("Received a POST request.\n"))
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", reqBody)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented) + "\n"))
	}
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8000", nil)
}

//curl 设置请求类型: curl -is -X POST/GET/... http://localhost:8000
