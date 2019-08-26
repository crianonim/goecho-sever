package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var r response
	fmt.Println(r)
	http.HandleFunc("/", logger(echoHandler))
	log.Fatal(http.ListenAndServe(":3000", nil))
}
