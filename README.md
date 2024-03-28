# Go Go Gadget Weather

This small API returns the weather condition given a latitude and longitude.

## Dependencies
 * go@1.22.1
 * chi@v5.0.12
 * [API Key from OpenWeatherMap](https://openweathermap.org/appid)
 * Docker (optional)
 * jq (optional)

## Running
### Natively
  1. Clone the repository
  2. `cd gogo-gadget-weather`
  3. `go mod download`
  4. `go run .`

### Docker
  1. Clone the repository
  2. `cd gogo-gadget-weather`
  3. `docker run -p 4242:4242 -ti $(docker build -q .)`

### Usage
Getting weather in Berlin
 ```bash
  curl -X GET \
    -H "X-API-KEY:$OPENWEATHERMAP_API_KEY" \
    http://localhost:4242/api/weather?lat=52.52&long=13.41 \
    | jq
 ```

 ```json
 {
  "weather": {
    "condition": "clear sky",
    "temperature": "moderate"
  }
}
 ```
