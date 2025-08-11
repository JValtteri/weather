# Open-Meteo Go API - weather/om
[![OM Build and Test](https://github.com/JValtteri/weather/actions/workflows/om-test.yaml/badge.svg)](https://github.com/JValtteri/weather/actions/workflows/om-test.yaml)

Package for interfacing with Open-Meteo APIs

### `om.Config()`
```go
func Config(key string, units string)
```
Config initializes the API. Run Config() before using the API.
`units` not implemented yet.
  - `metric` (default) Response unists are in Celsius, mm and m/s.
  - `imperial` Not implemented

### `om.Forecast()`
```go
func Forecast(lat float32, lon float32, model string, length string, units string) WeatherRange
```
Forecast fetches forecast data by coordinates and returns a WeatherRange object.
`lat` and `lon` are mandatory, other values can be set as empty string `""`. This will use the [stored defaults](#stored-defaults).

- `model string` The weather model to use. Defaults to `"best_model"`, which automatically chooses the model estimated to be the best for the selected region. For other options see [Open Meteo: Data Sources](https://open-meteo.com/en/docs?#data_sources)
- `length string` The length of requested forecast in days.
- `units string` Whether to use `metric` or `imperial` units. **(Not implemented)**. See: [`om.Config()`](#omconfig)

You can use `owm.Coord()` *(sic)* to find the coordinates of a cities.

### `om.Icon()`
```go
func Icon(id string, day bool) []byte
```
Icon returns the icon PNG image in byte form.

The `id` is a icon code that is found inside the forecast object
`day` is a bool value indicating whether to fetch a day or night version of the icon

### Configuring Default Parameters

These can be used to set **defaults**. Normally `om.Config()` sets these to useful defaults.

#### `om.API_CONFIG.API_KEY string`
API key for Open-Meteo premium features. **(Not implemented)**
Should be set through `om.Config()`.

#### `om.API_CONFIG.UNITS string`
Whether to use `metric` or `imperial` units. **(Not implemented)**
Can be set through `om.Config()`.

#### `om.API_CONFIG.MODEL string`
The weather model to use. Defaults to `"best_model"`, which automatically chooses the model estimated to be the best for the selected region.

#### `om.API_CONFIG.LENGTH string`
The length of requested forecast in days.

#### `om.API_CONFIG.NETWORK bool`
Setting to `false` disables networking for testing purposes.

`om.Config()` sets this to `true`. If you need networking disabled, set `om.API_CONFIG.NETWORK` to `false` **after** calling `om.Config()`, not before.

**Not use in production**

### Stored Defaults

| Key | Value
| :--: | :--: |
| `UNITS` | `"metric"` |
| `NETWORK` | `true` |
| `MODEL` | `"best_match"` |
| `LENGTH` | `"7"` |

### Response objects

### See: [`input_types.go`](input_types.go)

Some fields may not be implemented

---
For an example implementation, see [JValtteri/ll-weather-server](https://github.com/JValtteri/ll-weather-server)
