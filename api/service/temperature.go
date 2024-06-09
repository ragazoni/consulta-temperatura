package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Temperature struct {
	Celsius    float64 `json:"temp_c"`
	Fahrenheit float64 `json:"temp_f"`
	Kelvin     float64 `json:"temp_k"`
}

func GetTemperature(location string) (Temperature, error) {
	apiKey := ""
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, location)

	resp, err := http.Get(url)
	if err != nil {
		return Temperature{}, err
	}
	defer resp.Body.Close()

	var data struct {
		CurrentTemperatura struct {
			TempCelsius float64 `json:"temp_c"`
		} `json:"currentTemperatura"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return Temperature{}, err
	}

	celsius := data.CurrentTemperatura.TempCelsius
	fahrenheit := celsius*1.8 + 32
	kelvin := celsius + 273.15

	return Temperature{
		Celsius:    celsius,
		Fahrenheit: fahrenheit,
		Kelvin:     kelvin,
	}, nil

}
