package om

/* Open-Meteo API library OWM
 * Library for fetching forecasts using the
 * free Open-Meteo API
 * https://open-meteo.com/
 */

import (
    "net/http"
    "encoding/json"
    "fmt"
    "log"
    "strings"
    "io"
)

const MODEL = "best_match"

const FORECAST_URL string = "https://api.open-meteo.com/v1/forecast?latitude={LAT}&longitude={LON}&timezone=auto&hourly=temperature_2m,relative_humidity_2m,dew_point_2m,apparent_temperature,precipitation_probability,precipitation,rain,weather_code,pressure_msl,surface_pressure,cloud_cover,cloud_cover_low,cloud_cover_mid,cloud_cover_high,visibility,evapotranspiration,wind_speed_10m,wind_speed_80m,wind_direction_10m,wind_direction_80m,wind_gusts_10m,soil_temperature_0cm,soil_moisture_0_to_1cm,uv_index,is_day,sunshine_duration,wet_bulb_temperature_2m,boundary_layer_height,direct_radiation,diffuse_radiation&models={MODEL}&timeformat=unixtime&wind_speed_unit=ms"

/* Structure to store the configuration
 * Populate with owm.Config(key, units, countryCode string)
 */
type Api_config struct {
    API_KEY    string // Free version doesn't require a key
    UNITS      string // NOT IMPLEMENTED YET { (metric, imperial, standard (Kelvin) }
    COUNTRY    string // ISO 3166 country code
    NETWORK    bool   // Set to true. False disables requests
    MODEL      string // Weather model to use
}

var API_CONFIG Api_config

/* Run Config() first, to initialize the API
 */
func Config(key, units, countryCode string) {
    API_CONFIG.API_KEY    = key
    API_CONFIG.UNITS      = units
    API_CONFIG.COUNTRY    = countryCode
    API_CONFIG.NETWORK    = true
    API_CONFIG.MODEL      = "best_match"
}

/* Forecast fetches forecast data by coordinates and returns a WeatherRange object
 */
func Forecast(lat float32, lon float32) WeatherRange {
    var weather_obj WeatherRange
    var requestURL string = weatherURL(lat, lon, API_CONFIG.UNITS)
    var raw_weather []byte = request(requestURL)
    err := json.Unmarshal(raw_weather, &weather_obj)
    if err != nil {
        log.Println("Weather JSON Unmarshal error:", err)
    }
    return weather_obj
}

 // unexported

func weatherURL(lat, lon float32, units string) string {
    url := ""
    url = strings.Replace(FORECAST_URL, "{LAT}",     str_f(lat), 1)
    url = strings.Replace(url,          "{LON}",     str_f(lon), 1)
    url = strings.Replace(url,          "{UNITS}",   units, 1)
    url = strings.Replace(url,          "{API_KEY}", API_CONFIG.API_KEY, 1)
    url = strings.Replace(url,          "{MODEL}",   API_CONFIG.MODEL, 1)
    return url
}

// Make a request to chosen address
func request(address string) []byte {
    if !API_CONFIG.NETWORK {
        log.Println("Warning: API_CONFIG.NETWORK = false")
        return nil
    }
    resp, err := http.Get(address)
    if err != nil {
        log.Printf("Welp! GET from %s failed\n", address)
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Error reading response body:", err)
    }
    return body
}

// UTILS

func str_f(f float32) string {
    s := fmt.Sprintf("%f", f)
    return s
}
