package om

import (
)

/*
 * Structures for parsing raw input JSON data
 */

// Range Data //

type WeatherRange struct {
    Latitude                float32
    Longtitude              float32
    UTC_OFFSET_SECONDS      int
    Timezone                string
    Timezone_abbreviation   string
    Elevation               float32
    Hourly_units struct {
        Time                    string
        Temperature_2m          string
        Relative_humidity_2m    string
        Dew_point_2m            string
        Apparent_temperature    string
        Precipitation_probability   string
        Precipitation           string
        Rain                    string
        Weather_code            string
        pressure_msl            string
        Surface_pressure        string
        Cloud_cover             string
        Cloud_cover_low         string
        Cloud_cover_mid         string
        Cloud_cover_high        string
        Visibility              string
        Elevation               string
        Wind_speed_10m          string
        Wind_speed_80m          string
        Wind_direction_10m      string
        Wind_direction_80m      string
        Wind_gusts_10m          string
        Soil_temperature_0cm    string
        Soil_moisture_0_to_1cm  string
        Uv_index                string
        Is_day                  string
        Sunshine_duration       string
        Wet_bulb_temperature_2m string
        Boundary_layer_height   string
        Direct_radiation        string
        Diffuse_radiation       string
    }
    Hourly struct {
        Time                    []uint
        Temperature_2m          []float32
        Relative_humidity_2m    []int
        Dew_point_2m            []float32
        Apparent_temperature    []float32
        Precipitation_probability   []int
        Precipitation           []float32
        Rain                    []float32
        Weather_code            []int
        pressure_msl            []float32
        Surface_pressure        []float32
        Cloud_cover             []int
        Cloud_cover_low         []int
        Cloud_cover_mid         []int
        Cloud_cover_high        []int
        Visibility              []float32
        Elevation               []float32
        Wind_speed_10m          []float32
        Wind_speed_80m          []float32
        Wind_direction_10m      []int
        Wind_direction_80m      []int
        Wind_gusts_10m          []float32
        Soil_temperature_0cm    []float32
        Soil_moisture_0_to_1cm  []float32
        Uv_index                []float32
        Is_day                  []int
        Sunshine_duration       []float32
        Wet_bulb_temperature_2m []float32
        Boundary_layer_height   []float32
        Direct_radiation        []float32
        Diffuse_radiation       []float32
    }
    Daily struct {
        Time        []uint
        Sunrise     []uint
        Sunset      []uint
    }
}
