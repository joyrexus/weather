package weather

import (
	"net/http"
	"encoding/json"
)

type innerData struct { 
	Temp float64 `json:"temp"` 
}

type weatherData struct {
	City string    `json:"name"`
	Main innerData `json:"main"`
}

func Query(city string) (weatherData, error) {
	baseURL := "http://api.openweathermap.org/data/2.5/weather" 
	resp, err := http.Get(baseURL + "?q=" + city)
	if err != nil {
		return weatherData{}, err
	}
	defer resp.Body.Close()

	var data weatherData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	data.Main.Temp = kelvinToFahr(data.Main.Temp)

	return data, nil
}

func kelvinToFahr(k float64) float64 {
	return celsiusToFahr(k - 273.15)
}

func celsiusToFahr(c float64) float64 {
	return (c * 1.8) + 32
}
