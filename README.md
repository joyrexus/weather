# Weather Service Tutorial

What's this? Just working through Peter Bourgon's Go tutorial, [How I Start](https://howistart.org/posts/go/1), wherein he walks you through setting up a concurrent, [REST-style](https://github.com/rcc-uchicago/api-standards#rcc-web-service-standards) backend server for aggregating weather data from various weather service APIs using only the Go standard library. It's a nice intro to [Go-based web services](http://bl.ocks.org/joyrexus/e2daebf6759d6b930fe7).

See [`json-samples`](json-samples) for samples of the JSON-encoded data
returned by the [OpenWeatherMap](http://openweathermap.org/api) and 
[Weather Underground](http://www.wunderground.com/weather/api) APIs.

Note that the Weather Underground API requires an [API key](http://www.wunderground.com/weather/api/d/login.html). With `KEY` in hand, try ...

    curl http://api.wunderground.com/api/KEY/conditions/q/chicago.json

You can [test](http://api.openweathermap.org/data/2.5/weather?q=Chicago) the [OpenWeatherMap](http://openweathermap.org) API with ...

    curl http://api.openweathermap.org/data/2.5/weather?q=Chicago


## Setup

Install the weather package with `go get -u github.com/joyrexus/weather/...`.  This will also install the [`wx`](cmd/wx/README.md) and [`wxserve`](cmd/wxserve/README.md) commands.

Then, be sure to set the `WUNDERGROUND_API_KEY` environment variable to reflect your Weather Underground API key:

    export WUNDERGROUND_API_KEY=your_api_key

If `$GOPATH/bin` is in your `PATH` ...

    $ wx chicago
    35.5

    $ wx nashville
    65.1

... or run `wxserve` to start the demo web service and check the following urls in your browser:

    http://localhost:8080/weather/chicago

    http://localhost:8080/weather/nashville


