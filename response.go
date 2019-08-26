package main

import (
	"net/http"
	"net/url"
)

type response struct {
	Method    string
	Url       string
	Query     string
	Header    http.Header
	QueryList url.Values
}

type keyValue struct {
	key   string
	value string
}
