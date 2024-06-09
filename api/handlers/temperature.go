package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ragazoni/consulta-temperatura/api/service"
)

func GetTemperatureHandler(w http.ResponseWriter, r *http.Request) {
	zipcode := r.URL.Query().Get("zipcode")

	if len(zipcode) != 8 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	location, err := service.GetLocationByZipcode(zipcode)
	if err != nil {
		log.Printf("Error getting location: %+v", err)
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	temperature, err := service.GetTemperature(location)
	if err != nil {
		log.Printf("error getting temperature: %+v", err)
		http.Error(w, "failed to get temperature", http.StatusInternalServerError)
		return
	}

	resp := map[string]float64{
		"temp_C": temperature.Celsius,
		"temp_F": temperature.Fahrenheit,
		"temp_K": temperature.Kelvin,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
