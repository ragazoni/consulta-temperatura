package service

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Location struct {
	City string `json:"localidade"`
}

func GetLocationByZipcode(zipcode string) (string, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode)
	log.Printf("URL da solicitação: %s", url)

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var location Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return "", err
	}
	log.Printf("location.City: %s", location.City)
	return location.City, nil

}
