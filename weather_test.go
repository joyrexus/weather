package weather

import "testing"

func notWithin(b Bounds, x float64) bool {
	if x < b.min {
		return true
	}
	if x > b.max {
		return true
	}
	return false
}

type Bounds struct {
	min, max float64
}

var bounds = Bounds{-40.0, 130.0}

func TestOpenWeatherMap(t *testing.T) {
	w := openWeatherMap{}
	got, err := w.temperature("chicago")
	if err != nil {
		t.Error(err)
	}
	if notWithin(bounds, got) {
		t.Errorf("%d is not within normal temp range: %v", got, bounds)
	}
}

func TestWunderground(t *testing.T) {
	w := wunderground{}
	got, err := w.temperature("chicago")
	if err != nil {
		t.Error(err)
	}
	if notWithin(bounds, got) {
		t.Errorf("%d is not within normal temp range: %v", got, bounds)
	}
}

func TestWeatherServices(t *testing.T) {
	w := weatherServices{
		new(openWeatherMap),
		new(wundergroundMock),
	}
	got, err := w.temperature("chicago")
	if err != nil {
		t.Error(err)
	}
	if notWithin(bounds, got) {
		t.Errorf("%d is not within normal temp range: %v", got, bounds)
	}
}

func TestTemperature(t *testing.T) {
	got, err := Temperature("chicago")
	if err != nil {
		t.Fatal(err)
	}
	if notWithin(bounds, got) {
		t.Errorf("%d is not within normal temp range: %v", got, bounds)
	}
}
