package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	resObj := response{Method: r.Method, Url: r.URL.String(), Query: r.URL.RawQuery, QueryList: r.URL.Query(), Header: r.Header}
	if delay, _ := strconv.Atoi(r.URL.Query().Get("delay")); delay > 0 {
		time.Sleep(time.Duration(delay) * time.Second)
	}
	if status, _ := strconv.Atoi(r.URL.Query().Get("status")); status > 0 {
		w.WriteHeader(status)
		return
	}
	s, e := json.Marshal(resObj)
	if e != nil {
		w.WriteHeader(500)
		w.Write([]byte("Bad json"))
	}
	w.Write(s)

}
