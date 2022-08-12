// first website with go as backend
package main

import (
	f "fmt"
	h "net/http"
)

func main() {
	h.HandleFunc("/", func(w h.ResponseWriter, r *h.Request) {
		f.Println(w, "Hello, you have reuquested: %s\n", r.URL.Path)
	})

	h.ListenAndServe(":80", nil)
}
