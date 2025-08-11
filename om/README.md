# Open-Meteo Go API - weather/om
[![OM Build and Test](https://github.com/JValtteri/weather/actions/workflows/om-test.yaml/badge.svg)](https://github.com/JValtteri/weather/actions/workflows/om-test.yaml)

Package for interfacing with Open-Meteo APIs

### `om.Config()`
```go
func Config(key, units, model string, length string)
```
Config initializes the API. Run Config() before using the API.
`model` is the weather model to use. If left empty `""`, uses `"best_model"`, which automatically chooses the model estimated to be the best for the selected region.
`units` is not yet supported. Response unists are in Celsius, mm and m/s.
`length` is the requested forecast length in days.

### `om.Forecast()`
```go
func Forecast(lat float32, lon float32) WeatherRange
```
Forecast fetches forecast data by coordinates and returns a WeatherRange object.
Use Coord to find the coordinates of a cities.

### `om.Icon()`
```go
func Icon(id string, day bool) []byte
```
Icon returns the icon PNG image in byte form.

The `id` is a icon code that is found inside the forecast object
`day` is a bool value indicating whether to fetch a day or night version of the icon

### `om.API_CONFIG.NETWORK`
```go
owm.API_CONFIG.NETWORK bool
```
This variable is only used for testing. If you need to perform tests and want to disable networking, i.e. disable requests, set this variable to `false`.
```go
owm.API_CONFIG.NETWORK = false
```
`om.Config()` sets this to `true`. If you need networking disabled, set `om.API_CONFIG.NETWORK` to `false` **after** calling `om.Config()`, not before.

### Response objects

### See: [`input_types.go`](input_types.go)

Some fields may not be implemented

---
For an example implementation, see [JValtteri/ll-weather-server](https://github.com/JValtteri/ll-weather-server)
