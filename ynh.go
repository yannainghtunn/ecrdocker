package main

import (
	"fmt"
	"net/http"

)

func main() {
	fmt.Println("Hello, World")
	http.HandleFunc("/welcome", welcome)
	if err := http.ListenAndServe("0.0.0.0:9090", nil); err != nil {
		panic(err)
	}

}
func welcome(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome")
}
