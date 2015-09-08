package weather

import (
	"encoding/json"
	"io/ioutil"
)


/* W E A T H E R  S E R V I C E  M O C K S */


/* WEATHER UNDEGROUND SERVICE */

// wundergroundMock is a mock of the wunderground type
type wundergroundMock struct{               
	apiKey string
}

// temperature returns the temperature of `city` in fahrenheit.
func (w *wundergroundMock) temperature(city string) (float64, error) {
	path := PKG_DIR + "json-samples/wunderground.json"
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}

	var data struct {
		Main struct {
			Temp float64 `json:"temp_f"` 
		} `json:"current_observation"`
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return 0, err
	}

	return data.Main.Temp, nil
}


/* OPEN WEATHER MAP SERVICE */

// openWeatherMapMock is a mock of the openWeatherMap type.
type openWeatherMapMock struct{}

// temperature returns the temperature of `city` in fahrenheit.
func (w *openWeatherMapMock) temperature(city string) (float64, error) {
	path := PKG_DIR + "json-samples/openweathermap.json"
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}

	var data struct {
		Main struct {
			Temp float64 `json:"temp"` 
		} `json:"main"`
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return 0, err
	}

	return kelvinToFahr(data.Main.Temp), nil
}
