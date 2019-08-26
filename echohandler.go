package main

import (
	"fmt"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "echo")
}
