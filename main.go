package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"net/http"
    "net/url"

	"github.com/joho/godotenv"
)

type Request struct {
	Name string `json:"name"`
}

func getAccessToken(clientId string, clientSecret string, tokenEndpoint string) (string, error) {
	println(clientId, clientSecret, tokenEndpoint)

	data := url.Values{}
    data.Set("client_id", clientId)
    data.Set("client_secret", clientSecret)
    data.Set("scope", "https://graph.microsoft.com/.default")
    data.Set("grant_type", "client_credentials")

	resp, err := http.PostForm(tokenEndpoint, data)
    if err != nil {
        fmt.Println("Error making request:", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
    }

	return string(body), nil
}


func main() {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	clientId := os.Getenv("AZ_CLIENT_ID")
	clientSecret   := os.Getenv("AZ_CLIENT_SECRET")
	tokenEndpoint  := "https://login.microsoftonline.com/consumers/oauth2/v2.0/token"

    token, err := getAccessToken(clientId, clientSecret, tokenEndpoint)

    if err != nil {
        fmt.Println("Error getting access token:", err)
        return
    }

    fmt.Println("Access Token:", token)
}
