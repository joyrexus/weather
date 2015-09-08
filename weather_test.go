package weather

import "testing"

func TestOpenWeatherMap(t *testing.T) {
	w := openWeatherMapMock{}
	want := 75.65
	got, err := w.temperature("chicago")
	if err != nil {
		t.Error(err)
	}
	if got != want {
		t.Errorf("temp(%q) == %v, want %v", "chicago", got, want)
	}
}

func TestWunderground(t *testing.T) {
	w := wundergroundMock{}
	want := 78.8
	got, err := w.temperature("chicago")
	if err != nil {
		t.Error(err)
	}
	if got != want {
		t.Errorf("temp(%q) == %v, want %v", "chicago", got, want)
	}
}

func TestWeatherServices(t *testing.T) {
	w := weatherServices{
		&openWeatherMapMock{},
		&wundergroundMock{},
	}
	want := 77.225
	got, err := w.temperature("chicago")
	if err != nil {
		t.Error(err)
	}
	if got != want {
		t.Errorf("temp(%q) == %v, want %v", "chicago", got, want)
	}
}

func TestTemperature(t *testing.T) {
	want := 77.225
	got, err := Temperature("chicago")
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("temp(%q) == %v, want %v", "chicago", got, want)
	}
}
