package weather

import "os"

var PKG_DIR string = os.Getenv("GOPATH") + "/src/github.com/joyrexus/weather/"

// kelvingToFahr converts kelvin to fahrenheit.
func kelvinToFahr(k float64) float64 {
	return celsiusToFahr(k - 273.15)
}

// celsiusToFahr converts celsisus to fahrenheit.
func celsiusToFahr(c float64) float64 {
	return (c * 1.8) + 32
}

