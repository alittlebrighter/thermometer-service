// messy, messy, messy
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alittlebrighter/embd/sensor/mcp9808"
	nats "github.com/nats-io/nats.go"
)

var (
	pollInterval = "15s"
	natsURL      = nats.DefaultURL
	// stanCluster  string = "igor_cluster"
	// stanID       string = os.Getenv("OTTO_NODE")
)

func main() {
	flag.StringVar(&pollInterval, "pollFreq", pollInterval, "Frequency that the temperature sensor is checked for changes (duration string).")
	flag.StringVar(&natsURL, "natsUrl", natsURL, "Url for NATS instance to connect to.")
	// flag.StringVar(&stanID, "stanId", stanID, "The client ID used when connecting to streams.")
	flag.Parse()

	log.Println("Starting thermometer.")
	thermometer, err := NewThermometer()
	if err != nil {
		log.Fatalln("Error starting thermometer:", err)
	}
	defer thermometer.Shutdown()

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Println("Could not connect to message bus.")
	}

	serverApp := &server{}

	pollFreq, err := time.ParseDuration(pollInterval)
	if err != nil {
		pollFreq = 15 * time.Second
	}
	updates := make(chan interface{})
	timer := time.NewTicker(pollFreq)
	go poll(thermometer, timer, updates)
	go func(app *server, newTemps chan interface{}) {
		for temp := range newTemps {
			var toBroadcast interface{}

			switch temp.(type) {
			case error:
				app.Err = temp.(error)
				toBroadcast = app.Err
			case *mcp9808.Temperature:
				app.CurrentTemp = temp.(*mcp9808.Temperature)
				toBroadcast = convert(temp.(*mcp9808.Temperature))
			}

			broadcast(nc, toBroadcast)
		}
	}(serverApp, updates)

	path := "/temperature"

	http.HandleFunc(path, serverApp.ReadTemp)

	log.Println("Starting web server. Listening on port 80 at path: " + path)
	log.Fatal(http.ListenAndServe(":80", nil))
}

type server struct {
	CurrentTemp *mcp9808.Temperature
	Err         error
}

func (s *server) ReadTemp(w http.ResponseWriter, r *http.Request) {
	//tempC, err := thermometer.ReadTemperature()
	tempC := float64(-1)
	if s.CurrentTemp != nil {
		tempC = s.CurrentTemp.CelsiusDeg
	}
	fmt.Fprintf(w, fmt.Sprintf(`{"temperature":%f,"units":"Celsius","error":"%v"}`, tempC, s.Err))
}

func poll(thermometer *Thermometer, timer *time.Ticker, updates chan interface{}) {
	current, err := thermometer.ReadTemperature()

	if err != nil {
		log.Println("could not read from thermometer:", err)
		return
	}
	updates <- current

	for {
		<-timer.C
		newTemp, err := thermometer.ReadTemperature()
		if err != nil {
			updates <- current
			continue
		} else if newTemp.CelsiusDeg == current.CelsiusDeg {
			continue
		}

		current = newTemp
		updates <- current
	}
}

const topic = "otto.sensor.temperature.current"

// Broadcaster interface with this definition probably makes sense (without the connection argument)
func broadcast(conn *nats.Conn, value interface{}) {
	data, _ := json.Marshal(&SensorUpdate{Location: os.Getenv("OTTO_LOCATION"), Type: "ambient temperature", Value: value})

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
