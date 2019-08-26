package main

import (
	"encoding/json"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	resObj := response{Method: r.Method, Url: r.URL.String(), Query: r.URL.RawQuery, QueryList: r.URL.Query(), Header: r.Header}
	// fmt.Println(resObj, r.URL.Query(), r.Header)
	s, e := json.Marshal(resObj)
	if e != nil {
		w.WriteHeader(500)
		w.Write([]byte("Bad json"))

	}
	w.Write(s)

}
