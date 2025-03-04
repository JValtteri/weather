# OpenWeatherMap Go API - weather/owm

Package for interfacing with openweathermap.org basic APIs

### `owm.Config()`
```go
func Config(key, units, countryCode string, network bool)
```
Config initializes the API. Run Config() before using the API. Set network to true.

### `owm.Forecast()`
```go
func Forecast(lat float32, lon float32) WeatherRange
```
Forecast fetches forecast data by coordinates and returns a WeatherRange object.
Use Coord to find the coordinates of a cities.

### `owm.Coord()`
```go
func Coord(name string) (float32, float32)
```
Coord returns the coordinates for the given city

### `owm.Icon()`
```go
func Icon(id string) []byte
```
Icon returns the icon PNG image in byte form.

The `id` is a icon code that is found inside the forecast object

### `owm.API_CONFIG.NETWORK`
```go
owm.API_CONFIG.NETWORK bool
```
This variable is only used for testing. If you need to perform tests and want to disable networking, i.e. disable requests, set this variable to `false`.
```go
owm.API_CONFIG.NETWORK = false
```
`owm.Config()` sets this to `true`. If you need networking disabled, set `owm.API_CONFIG.NETWORK` to `false` **after** calling `owm.Config()`, not before.

### Response objects

### See: [`input_types.go`](input_types.go)

Some fields may not be implemented

---
For an example implementation, see [JValtteri/ll-weather-server](https://github.com/JValtteri/ll-weather-server)
