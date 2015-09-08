package weather

import (
	"net/http"
	"encoding/json"
)

// Temperature returns the temperature of `city` in fahrenheit.
func Temperature(city string) (float64, error) {
	w := weatherServices{
		new(wunderground),
		new(openWeatherMap),
	}
	t, err := w.temperature(city)
	if err != nil {
		return 0, err
	}
	return t, nil
}

type weatherService interface {
	temperature(city string) (float64, error) // in Fahrenheit!
}

type weatherServices []weatherService

func (ws weatherServices) temperature(city string) (float64, error) {
	sum := 0.0
	
	for _, service := range ws {
		t, err := service.temperature(city)
		if err != nil {
			return 0, err
		}
		sum += t
	}

	return sum / float64(len(ws)), nil
}


/* W E A T H E R  S E R V I C E S */


/* WEATHER UNDEGROUND SERVICE */

// wunderground is a weatherService that utilizes the WeatherUnderground API.
type wunderground struct{}               

// temperature returns the temperature of `city` in fahrenheit.
func (w *wunderground) temperature(city string) (float64, error) {
	err := checkKey() // ensure that API_KEY is set!
	if err != nil {
		return 0, err
	}
	baseURL := "http://api.wunderground.com/api/" + API_KEY + "/conditions/q/" 
	resp, err := http.Get(baseURL + city + ".json")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data struct {
		Main struct {
			Temp float64 `json:"temp_f"` 
		} `json:"current_observation"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	return data.Main.Temp, nil
}


/* OPEN WEATHER MAP SERVICE */

// openWeatherMap is a weatherService that utilizes the OpenWeatherMap API.
type openWeatherMap struct{}

// temperature returns the temperature of `city` in fahrenheit.
func (w *openWeatherMap) temperature(city string) (float64, error) {
	baseURL := "http://api.openweathermap.org/data/2.5/weather" 
	resp, err := http.Get(baseURL + "?q=" + city)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data struct {
		Main struct {
			Temp float64 `json:"temp"` 
		} `json:"main"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	return kelvinToFahr(data.Main.Temp), nil
}
