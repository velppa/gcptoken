package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	token := token()
	bytes, err := json.Marshal(token)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Print(string(bytes))
}

func token() *oauth2.Token {
	credentialsFileName := os.Getenv("GCP_ACCOUNT_CREDENTIALS")
	b, err := ioutil.ReadFile(credentialsFileName)
	if err != nil {
		log.Fatalf("Reading %s failed: %s", credentialsFileName, err)
	}
	scopes := strings.Split(os.Getenv("GCP_SCOPE"), ",")
	conf, err := google.JWTConfigFromJSON(b, scopes...)
	if err != nil {
		log.Fatalf("Reading config from file %s failed: %s", credentialsFileName, err)
	}
	token, err := conf.TokenSource(oauth2.NoContext).Token()
	if err != nil {
		log.Fatalf("Token source failed: %s", err)
	}
	return token
}
