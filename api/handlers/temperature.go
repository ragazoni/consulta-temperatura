package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/ragazoni/consulta-temperatura/api/service"
)

func GetTemperatureHandler(w http.ResponseWriter, r *http.Request) {
	zipcode := r.URL.Query().Get("zipcode")

	if !isValidCEP(zipcode) {
		http.Error(w, `{"message": "invalid zipcode"}`, http.StatusUnprocessableEntity)
		return
	}

	location, err := service.GetLocationByZipcode(zipcode)
	log.Printf("location: %+v", location)
	if err != nil {
		http.Error(w, `{"message": "can not find zipcode"}`, http.StatusNotFound)
		return
	}

	tempC, err := service.GetTemperature(location)
	log.Printf("tempC: %+v", tempC)
	if err != nil {
		http.Error(w, `{"message": "unable to fetch weather"}`, http.StatusInternalServerError)
		return
	}

	tempF := tempC*1.8 + 32
	tempK := tempC + 273

	response := map[string]float64{
		"temp_C": tempC,
		"temp_F": tempF,
		"temp_K": tempK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func isValidCEP(zipcode string) bool {
	match, _ := regexp.MatchString(`^\d{8}$`, zipcode)
	return match
}
