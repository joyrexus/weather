package weather

import "testing"

func TestOpenWeatherMapMock(t *testing.T) {
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

func TestWundergroundMock(t *testing.T) {
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

func TestWeatherServicesMock(t *testing.T) {
	w := weatherServices{
		new(openWeatherMapMock),
		new(wundergroundMock),
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
