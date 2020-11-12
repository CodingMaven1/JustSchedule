package main

import (
	"log"

	"github.com/gorilla/mux"
)

// AddRoutes for adding the routes to the router...
func AddRoutes(router *mux.Router) {
	log.Println("Loading the routes...")

	router.HandleFunc("/", RenderHome)

	log.Println("Routes Loaded!")
}
