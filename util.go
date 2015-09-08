package weather

import (
	"os"
	"errors"
)

var PKG_DIR string = os.Getenv("GOPATH") + "/src/github.com/joyrexus/weather/"

var API_KEY string = os.Getenv("WUNDERGROUND_API_KEY")

func checkKey() error {
	msg := "The WUNDERGROUND_API_KEY environment variable should contain "
	msg += "your weather underground API key!"
	if API_KEY == "" {
		return errors.New(msg)
	}
	return nil
}

// kelvingToFahr converts kelvin to fahrenheit.
func kelvinToFahr(k float64) float64 {
	return celsiusToFahr(k - 273.15)
}

// celsiusToFahr converts celsisus to fahrenheit.
func celsiusToFahr(c float64) float64 {
	return (c * 1.8) + 32
}

