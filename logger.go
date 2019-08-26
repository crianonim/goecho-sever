package main

import (
	"fmt"
	"net/http"
	"time"
)

func logger(h func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Log", time.Now().UTC(), r.Method, r.URL)
		h(w, r)
	})
}
