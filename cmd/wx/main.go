package main

import (
	"os"
	"log"
	"fmt"

	"github.com/joyrexus/weather"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please specify a city name!")
	}
	city := os.Args[1]
	temp, err := weather.Temperature(city)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(temp)
}
