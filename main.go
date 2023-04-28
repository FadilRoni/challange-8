package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type WeatherValues struct {
	ValueWater int `json:"value_water"`
	ValueWind  int `json:"value_wind"`
}

type WeatherData struct {
	Data        WeatherValues `json:"data"`
	StatusWater string        `json:"status_water"`
	StatusWind  string        `json:"status_wind"`
}

func main() {

	timer := time.NewTimer(1 * time.Minute)
	ticker := time.NewTicker(15 * time.Second)

	for {
		select {
		case <-ticker.C:
			valueWater := rand.Intn(100)
			valueWind := rand.Intn(100)
			statusWater := getStatusWater(valueWater)
			statusWind := getStatusWind(valueWind)

			weatherData := WeatherData{
				Data: WeatherValues{
					ValueWater: valueWater,
					ValueWind:  valueWind,
				},
				StatusWater: statusWater,
				StatusWind:  statusWind,
			}

			payload, _ := json.Marshal(weatherData)
			_, err := http.Post("http://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(payload))
			if err != nil {
				return
			}
			fmt.Printf("{\n water: %v, \n wind: %v\n}\n status water: %v\n status wind: %v\n", weatherData.Data.ValueWater, weatherData.Data.ValueWind, weatherData.StatusWater, weatherData.StatusWind)

		case <-timer.C:
			log.Fatalln("Selesai")
			return
		}
	}
}

func getStatusWater(Water int) string {
	if Water < 5 {
		return "aman"
	} else if Water >= 5 && Water <= 8 {
		return "siaga"
	} else {
		return "bahaya"
	}
}

func getStatusWind(Wind int) string {
	if Wind < 6 {
		return "aman"
	} else if Wind >= 7 && Wind <= 15 {
		return "siaga"
	} else {
		return "bahaya"
	}
}
