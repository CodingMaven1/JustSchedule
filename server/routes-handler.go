package main

import (
	"io/ioutil"
	"context"
	"net/http"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}

type Response struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

// Signup for registering a user...
func Signup(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "applications/json")

	var user User
	var result User
	var response Response

	body, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.Error = "There was some error!"
		json.NewEncoder(res).Encode(response)
		return
	}
	users, err := ConnectToDB()
	if err != nil {
		response.Error = "Some DB error!"
		json.NewEncoder(res).Encode(response)
		return
	}

	err = users.FindOne(context.Background(), bson.D{{"username", user.Username}}).Decode(&result)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

			if err != nil {
				response.Error = "There was some error while hashing passwords."
				json.NewEncoder(res).Encode(response)
				return
			}
			user.Password = string(hash)
			_, err = users.InsertOne(context.Background(), user)

			if err != nil{
				response.Error = "There was some error while creating account!"
				json.NewEncoder(res).Encode(response)
				return
			}

			response.Result = "Registration Successfull"
			json.NewEncoder(res).Encode(response)
			return
		}

		response.Error = err.Error()
		json.NewEncoder(res).Encode(response)
		return
	}

	response.Result = "Username already exists!"
	json.NewEncoder(res).Encode(response)
	return
}

func Login(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var user User
	var result User
	var response Response

	body, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &user)

	if err != nil {
		response.Error = "There was some error!"
		json.NewEncoder(res).Encode(response)
		return
	}

	users, err := ConnectToDB()

	if err != nil{
		response.Error = "Error while connecting to DB"
		json.NewEncoder(res).Encode(response)
		return
	}

	 err = users.FindOne(context.Background(), bson.D{{"username", user.Username}}).Decode(&result)

	 if err != nil {
		 response.Error = "Invalid Username"
		 json.NewEncoder(res).Encode(response)
		 return
	 }

	 err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))

	 if err != nil{
		 response.Error = "Wrong Password"
		 json.NewEncoder(res).Encode(response)
		 return
	 }

	 token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		 "username": result.Username,
		 "firstname": result.Firstname,
		 "lastname": result.Lastname,
	 })

	 tokenString, err := token.SignedString([]byte("JustSchedule"))

	 if err != nil {
		 response.Error = "Error while generating token"
		 json.NewEncoder(res).Encode(response)
		 return
	 }

	 result.Password = ""
	 result.Token = tokenString
	 json.NewEncoder(res).Encode(result)
	 return
}