package main

import (
	"net/http"
)

// Signup for registering a user...
func Signup(res http.ResponseWriter, req *http.Request) {
	var resp []byte
	res.WriteHeader(http.StatusOK)
	res.Write(resp)
}
