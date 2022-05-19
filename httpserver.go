package main

import (
	"fmt"
	"net/http"
)

func hhhh(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hhhh \n ")
}

func main() {
	http.HandleFunc("/hhhh", hhhh)
	http.ListenAndServe(":8080", nil)
}
