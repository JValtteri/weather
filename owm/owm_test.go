package owm

import (
    "testing"
)

const TEST_URL string = "localhost:55005?lat={LAT}&lon={LON}&units={UNITS}&q={CITY_NAME},{COUNTRY_CODE}&limit={LIMIT}"

const EXP_WURL string = "https://api.openweathermap.org/data/2.5/forecast?lat=12.300000&lon=4.560000&units=metric&appid=test"
const EXP_CURL string = "https://api.openweathermap.org/geo/1.0/direct?q=atlantis,NN&limit=1&appid=test"
const EXP_IURL string = "https://openweathermap.org/img/wn/NN@2x.png"

func setup() {
    Config("test", "metric", "NN", false)
}

func TestRequestDisable(t *testing.T) {
    if request(TEST_URL) != nil {
        t.Errorf("Network disable failed: %v", API_CONFIG.NETWORK)
    }
}

func TestWeatherUrl(t *testing.T) {
    setup()
    url := weatherURL(12.3, 4.56, "metric")
    if url != EXP_WURL {
        t.Errorf("URL format failed %s", url)
        t.Errorf("Expected URL      %s", EXP_WURL)
    }
}

func TestCityUrl(t *testing.T) {
    setup()
    url := cityURL("atlantis", "NN", "1")
    if url != EXP_CURL {
        t.Errorf("URL format failed %s", url)
        t.Errorf("Expected URL      %s", EXP_CURL)
    }
}

func TestIconUrl(t *testing.T) {
    setup()
    url := iconURL("NN")
    if url != EXP_IURL {
        t.Errorf("URL format failed %s", url)
        t.Errorf("Expected URL      %s", EXP_IURL)
    }
}

func TestStrF(t *testing.T) {
    exp := "3.141500"
    if str_f(3.1415) != exp {
        t.Errorf("String did not match: %v : %v", exp, str_f(3.1415))
    }
}

func TestCoord(t *testing.T) {
    city := InCity{{
        Name:   "atlantis",
        Lat:     3.14,
        Lon:     14.5,
        Country: "NN",
    }}
    lat, lon := coord(city)
    if lat != 3.14 || lon != 14.5 {
        t.Errorf("City coord didn't match %v %v", lat, lon)
    }
}
