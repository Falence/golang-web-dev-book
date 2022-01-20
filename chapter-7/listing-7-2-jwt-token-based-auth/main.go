package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// Using asymmetric crypto/RSA keys
// Location of the files used for signing and verification
const (
	privateKeyPath = "keys/app.rsa"	// openssl genrsa -out app.rsa 1024
	publicKeyPath = "keys/app.rsa.pub"	// openssl rsa -in app.rsa -pubout > app.rsa.pub
)

// Verify key and sign key
var (
	verifyKey, signKey []byte
)

// Struct User for parsing login credentials
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Read the key files before starting http handlers
func init() {
	var err error

	signKey, err = ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	verifyKey, err = ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}
}

// Reads the login credentials, checks them and creates the JWT token
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	// Decode into User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error in request body")
		return
	}

	// Validate user credentials
	if user.UserName != "falence" && user.Password != "pass" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Wrong info")
		return
	}

	// Create a signer for rsa 256
	t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), jwt.MapClaims{
		"iss": "admin",
		"CustomUserInfo": struct {
			Name string
			Role string
		}{user.UserName, "Member"},
		"exp": time.Now().Add(time.Minute * 20).Unix(),
	})

	// // Set our claims
	// t.Claims["iss"] = "admin"
	// t.Claims["CustomUserInfo"] = struct {
	// 	Name string
	// 	Role string
	// }{user.UserName, "Member"}

	// // Set the expiration time
	// t.Claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	tokenString, err := t.SignedString(signKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Sorry, error while signing token!")
		log.Printf("Token signing error: %v\n", err)
		return
	}

	// Respond
	response := Token{tokenString}
	jsonResponse(response, w)
} 

// Only accessible with valid token
func authHandler(w http.ResponseWriter, r *http.Request) {
	// validate the token
	token, err := jwt.Parse
}