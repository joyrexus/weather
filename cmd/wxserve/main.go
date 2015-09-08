package main

import (
    "net/http"
    "strings"
	"encoding/json"

	"github.com/joyrexus/weather"
)

func main() {
    http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
        city := strings.SplitN(r.URL.Path, "/", 3)[2]

        temp, err := weather.Temperature(city)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json; charset=utf-8")
		data := map[string]interface{}{
			"city": city,
			"temp": temp,
		}
        json.NewEncoder(w).Encode(data)
    })

    http.ListenAndServe(":8080", nil)
}
