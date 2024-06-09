package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	City string `json:"localidade"`
}

func GetLocationByZipcode(zipcode string) (string, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var location Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return "", err
	}

	return location.City, nil

}
