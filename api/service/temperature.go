package service

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Temperature struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetTemperature(location string) (float64, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	apiKey := os.Getenv("WEATHER_API_KEY")
	encodedLocation := url.QueryEscape(location)
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?q=%s&key=%s", encodedLocation, apiKey)
	log.Printf("URL da solicitação: %s", url)

	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return 0, fmt.Errorf("location not found")
	} else if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("error fetching weather")
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	current, ok := data["current"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("unable to parse response")
	}

	tempC, ok := current["temp_c"].(float64)
	if !ok {
		return 0, fmt.Errorf("temperature not found in response")
	}

	return tempC, nil
}
