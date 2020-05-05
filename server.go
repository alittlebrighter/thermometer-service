package main

import (
	"fmt"
	"log"
	"net/http"
)

func TempCToF(tempC float64) float64 {
	return tempC*9/5 + 32
}

func TempFToC(tempF float64) float64 {
	return (tempF - 32) * 5 / 9
}

func main() {

	log.Println("Starting thermometer.")
	thermometer, err := NewThermometer()
	if err != nil {
		log.Fatalln("Error starting thermometer: " + err.Error())
	}
	defer thermometer.Shutdown()

	path := "/temperature"

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		tempC, err := thermometer.ReadTemperature()
		fmt.Fprintf(w, fmt.Sprintf(`{"temperature":%f,"units":"Celsius","error":"%v"}`, tempC.CelsiusDeg, err))
	})

	log.Println("Starting web server. Listening on port 80 at path: " + path)
	log.Fatal(http.ListenAndServe(":80", nil))
}
