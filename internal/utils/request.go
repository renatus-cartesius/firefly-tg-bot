package utils

import (
	"io"
	"log"
	"net/http"
)

type AuthClient struct {
	token      string
	httpClient *http.Client
}

func NewAuthClient(token string) *AuthClient {
	return &AuthClient{
		token:      token,
		httpClient: &http.Client{},
	}
}

func (ac *AuthClient) Get(url string) (string, error) {
	bearer := "Bearer " + ac.token
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)

	resp, err := ac.httpClient.Do(req)
	if err != nil {
		log.Fatalln("Error on response: ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error on reading response body: ", err)
	}

	return string([]byte(body)), nil
}
