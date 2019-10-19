//HTTP Server
package main

import (
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//Accept 报头: text/plain、text/html、application/json、application/xml
	switch r.Header.Get("Accept") {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte(`{"Message": "Hello world"}`))
	default:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("Hello world"))
	}

}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8000", nil)
}

//curl 设置报头: curl -is -H 'Accept: application/json' http://localhost:8000
