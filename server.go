package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alittlebrighter/embd/sensor/mcp9808"
	jsoniter "github.com/json-iterator/go"
	nats "github.com/nats-io/nats.go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	log.Println("Starting thermometer.")
	thermometer, err := NewThermometer()
	if err != nil {
		log.Fatalln("Error starting thermometer: " + err.Error())
	}
	defer thermometer.Shutdown()

	nc, _ := nats.Connect(nats.DefaultURL)

	serverApp := &server{}

	updates := make(chan *mcp9808.Temperature)
	timer := time.NewTicker(15 * time.Second)
	go poll(thermometer, timer, updates)
	go func(app *server, newTemps chan *mcp9808.Temperature) {
		for temp := range newTemps {
			app.CurrentTemp = temp

			converted := convert(temp)
			broadcast(nc, converted)
		}
	}(serverApp, updates)

	path := "/temperature"

	http.HandleFunc(path, serverApp.ReadTemp)

	log.Println("Starting web server. Listening on port 80 at path: " + path)
	log.Fatal(http.ListenAndServe(":80", nil))
}

type server struct {
	CurrentTemp *mcp9808.Temperature
}

func (s *server) ReadTemp(w http.ResponseWriter, r *http.Request) {
	//tempC, err := thermometer.ReadTemperature()
	fmt.Fprintf(w, fmt.Sprintf(`{"temperature":%f,"units":"Celsius","error":"%v"}`, s.CurrentTemp.CelsiusDeg, nil))
}

func poll(thermometer *Thermometer, timer *time.Ticker, updates chan *mcp9808.Temperature) {
	current, err := thermometer.ReadTemperature()
	if err != nil {
		log.Println("could not read from thermometer:", err)
		return
	}
	updates <- current

	for {
		<-timer.C
		newTemp, err := thermometer.ReadTemperature()
		if err != nil || newTemp.CelsiusDeg == current.CelsiusDeg {
			continue
		}

		current = newTemp
		updates <- current
	}
}

const topic = "igor.sensor.temperature"

// Broadcaster interface with this definition probably makes sense (without the connection argument)
func broadcast(conn *nats.Conn, value interface{}) {
	data, _ := json.Marshal(&SensorUpdate{Location: "living room", Type: "ambient temperature", Value: value})

	conn.Publish(topic, data)
}

func convert(sensorTemp *mcp9808.Temperature) *Temperature {
	return &Temperature{
		Degrees: sensorTemp.CelsiusDeg,
		Unit:    Celsius,
	}
}

type SensorUpdate struct {
	Location string      `json:"location"`
	Type     string      `json:"type"`
	Value    interface{} `json:"value"`
}

type Temperature struct {
	Degrees float64         `json:"degrees"`
	Unit    TemperatureUnit `json:"unit"`
}

type TemperatureUnit string

const (
	Celsius    TemperatureUnit = "Celsius"
	Fahrenheit TemperatureUnit = "Fahrenheit"
)

func TempCToF(tempC float64) float64 {
	return tempC*9/5 + 32
}

func TempFToC(tempF float64) float64 {
	return (tempF - 32) * 5 / 9
}
