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
func timer(name string) {
	time.Sleep(1000)
	fmt.Println("Name is " + name)
}
func welcome(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome")
}
