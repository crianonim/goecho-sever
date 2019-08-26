package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", logger(echoHandler))
	log.Fatal(http.ListenAndServe(":3000", nil))
}
