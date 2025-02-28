package owm

/* OpenWeatherMap API library OWM
 * Library for fetching forecasts using the
 * free OpenWeatherMap API
 */

import (
    "net/http"
    "encoding/json"
    "fmt"
    "log"
    "strings"
    "io"
)

const FORECAST_URL string = "https://api.openweathermap.org/data/2.5/forecast?lat={LAT}&lon={LON}&units={UNITS}&appid={API_KEY}"
const CITY_URL string = "https://api.openweathermap.org/geo/1.0/direct?q={CITY_NAME},{COUNTRY_CODE}&limit={LIMIT}&appid={API_KEY}"
const ICON_URL = "https://openweathermap.org/img/wn/{ICON}@2x.png"

/* Structure to store the configuration
 * Populate with owm.Config(key, units, countryCode string)
 */
type Api_config struct {
    API_KEY    string
    UNITS      string // metric, imperial, standard (Kelvin)
    COUNTRY    string // ISO 3166 country code
    CITY_LIMIT string
}

var API_CONFIG Api_config

/* Run Config() first, to initialize the API
 */
func Config(key, units, countryCode string) {
    API_CONFIG.API_KEY    = key
    API_CONFIG.UNITS      = units
    API_CONFIG.COUNTRY    = countryCode
    API_CONFIG.CITY_LIMIT = "1"
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

/* Coord returns the coordinates for the given city
 */
func Coord(name string) (float32, float32) {
    var city_obj InCity
    var requestURL string = cityURL(name, API_CONFIG.COUNTRY, API_CONFIG.CITY_LIMIT)
    var raw_city []byte   = request(requestURL)
    unmarshalCity(raw_city, &city_obj)
    if len(city_obj) == 0 {
        return 0.0, 0.0
    }
    var lat, lon float32 = coord(city_obj)
    return lat, lon
}

/* Icon returns the icon image in byte form.
 */
func Icon(id string) []byte {
    var requestURL string = makeIconURL(id)
    var raw_icon []byte   = request(requestURL)
    return raw_icon
}

 // unexported

func unmarshalCity(raw_city []byte, city_obj *InCity) {
    err := json.Unmarshal(raw_city, city_obj)
    if err != nil {
        log.Println("City JSON Unmarshal error:", err)
    }
}

func coord(city_obj InCity) (float32, float32) {
    var lat float32 = city_obj[0].Lat
    var lon float32 = city_obj[0].Lon
    return lat, lon
}

func weatherURL(lat, lon float32, units string) string {
    url := ""
    url = strings.Replace(FORECAST_URL, "{LAT}",     str_f(lat), 1)
    url = strings.Replace(url,          "{LON}",     str_f(lon), 1)
    url = strings.Replace(url,          "{UNITS}",   units, 1)
    url = strings.Replace(url,          "{API_KEY}", API_CONFIG.API_KEY, 1)
    return url
}

func cityURL(name, country, limit string) string {
    url := ""
    url = strings.Replace(CITY_URL, "{CITY_NAME}",    name, 1)
    url = strings.Replace(url,      "{COUNTRY_CODE}", country, 1)
    url = strings.Replace(url,      "{LIMIT}",        limit, 1)
    url = strings.Replace(url,      "{API_KEY}",      API_CONFIG.API_KEY, 1)
    return url
}

func makeIconURL(id string) string {
    url := ""
    url = strings.Replace(ICON_URL, "{ICON}", id, 1)
    return url
}

// Make a request to chosen address
func request(address string) []byte {
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
