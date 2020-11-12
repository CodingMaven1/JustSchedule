package main

import (
	"net/http"
)

// RenderHome for rendering the landing page...
func RenderHome(res http.ResponseWriter, req *http.Request) {
	var resp []byte
	res.WriteHeader(http.StatusOK)
	res.Write(resp)
}
