package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()
	AddRoutes(router)
	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
