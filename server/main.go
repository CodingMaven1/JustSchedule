package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server is running on port 8000")
	var router = mux.NewRouter()
	AddRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
