package weather

import "testing"

func TestWeatherData(t *testing.T) {
	type bounds struct {
		min float64
		max float64
	}
	type expected struct {
		city string
		bounds
	}

	cases := []struct {
		in string
		want expected
	}{
		{
			"chicago", 
			expected{"Chicago", bounds{-50.00, 150.00}},
		},
		{
			"nashville", 
			expected{"Nashville", bounds{-50.00, 150.00}},
		},
	}

	for _, c := range cases {
		got, err := Query(c.in)
		if err != nil {
			t.Error(err)
		}
		if got.City != c.want.city {
			t.Errorf("query(%q) == %q, want %q", c.in, got.City, c.want.city)
		}
		if got.Main.Temp > c.want.bounds.max {
			t.Errorf(
				"%v fahrenheit %q is unusually hot!", 
				got.City, 
				got.Main.Temp,
			)
		}
		if got.Main.Temp < c.want.bounds.min {
			t.Errorf(
				"%v fahrenheit for %q is unusually cold!", 
				got.City, 
				got.Main.Temp,
			)
		}
	}
}
