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
