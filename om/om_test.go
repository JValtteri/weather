package om

import (
    "testing"
)

const TEST_URL string = "localhost:55005?lat={LAT}&lon={LON}&units={UNITS}&q={CITY_NAME},{COUNTRY_CODE}&limit={LIMIT}"
const EXP_WURL string = "https://api.open-meteo.com/v1/forecast?latitude=12.300000&longitude=4.560000&timezone=auto&hourly=temperature_2m,relative_humidity_2m,dew_point_2m,apparent_temperature,precipitation_probability,precipitation,rain,weather_code,pressure_msl,surface_pressure,cloud_cover,cloud_cover_low,cloud_cover_mid,cloud_cover_high,visibility,evapotranspiration,wind_speed_10m,wind_speed_80m,wind_direction_10m,wind_direction_80m,wind_gusts_10m,soil_temperature_0cm,soil_moisture_0_to_1cm,uv_index,is_day,sunshine_duration,wet_bulb_temperature_2m,boundary_layer_height,direct_radiation,diffuse_radiation&models=best_match&forecast_days=14&timeformat=unixtime&wind_speed_unit=ms"
const EXP2_WURL string = "https://api.open-meteo.com/v1/forecast?latitude=12.300000&longitude=4.560000&timezone=auto&hourly=temperature_2m,relative_humidity_2m,dew_point_2m,apparent_temperature,precipitation_probability,precipitation,rain,weather_code,pressure_msl,surface_pressure,cloud_cover,cloud_cover_low,cloud_cover_mid,cloud_cover_high,visibility,evapotranspiration,wind_speed_10m,wind_speed_80m,wind_direction_10m,wind_direction_80m,wind_gusts_10m,soil_temperature_0cm,soil_moisture_0_to_1cm,uv_index,is_day,sunshine_duration,wet_bulb_temperature_2m,boundary_layer_height,direct_radiation,diffuse_radiation&models=NN&forecast_days=7&timeformat=unixtime&wind_speed_unit=ms"

func setup() {
    Config("test", "metric")
    API_CONFIG.LENGTH = "14"
    API_CONFIG.NETWORK = false
}

func TestRequestDisable(t *testing.T) {
    if request(TEST_URL) != nil {
        t.Fatalf("Network disable failed: %v", API_CONFIG.NETWORK)
    }
}

func TestDefaultAssign(t *testing.T) {
    expect := "bar"
    out := defaults("foo", expect)
    if out != "bar" {
        t.Fatalf("Expected: %s, got: %s\n", expect, out)
    }
}

func TestDefaultEmpty(t *testing.T) {
    expect := "bar"
    out := defaults(expect, "")
    if out != "bar" {
        t.Fatalf("Expected: %s, got: %s\n", expect, out)
    }
}

func TestWeatherUrl(t *testing.T) {
    setup()
    url := weatherURL(12.3, 4.56, "", "", "metric")
    if url != EXP_WURL {
        t.Errorf("URL format failed %s", url)
        t.Errorf("Expected URL      %s", EXP_WURL)
    }
}

func TestAltWeatherUrl(t *testing.T) {
    Config("test", "metric")
    API_CONFIG.MODEL="NN"
    API_CONFIG.NETWORK = false
    url := weatherURL(12.3, 4.56, "", "", "metric")
    if url != EXP2_WURL {
        t.Errorf("URL format failed %s", url)
        t.Errorf("Expected URL      %s", EXP2_WURL)
    }
}

func TestStrF(t *testing.T) {
    exp := "3.141500"
    if str_f(3.1415) != exp {
        t.Errorf("String did not match: %v : %v", exp, str_f(3.1415))
    }
}
