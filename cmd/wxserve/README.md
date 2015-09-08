# wxserve

A simple http server that provides a JSON-encoded response indicating the current temperature for a particular city.  

The city name should be specified as the last component of `http://localhost:8080/weather/CITY`.  For example, to get weather data for the city of Nashville, start the server (`go run` in this directory) and try ...

    curl http://localhost:8080/weather/nashville

... or install with `go install github.com/joyrexus/weather/cmd/wxserve` and then try ...

    $GOPATH/bin/wxserve

... to start the server.
