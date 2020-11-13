package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

//Client to access database...
var Client *mongo.Client

func main() {
	Client, err := ConnectToDB()
	if err != nil {
		log.Fatal(err, Client)
	}
	var router = mux.NewRouter()
	AddRoutes(router)
	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
