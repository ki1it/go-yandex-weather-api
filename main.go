package main

import (
	"fmt"
	weather "github.com/3crabs/go-yandex-weather-api"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"strconv"
)

type InfoResponse struct {
	CurrTemp string
}
type Error struct {
	Error string
}

type Request struct {
	Lat string
	Lon string
}

func getEnv(key string, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return val
}

func main() {
	_ = godotenv.Load()
	log.Println("App started")
	yandexWeatherApiKey := getEnv("YW_API_KEY", "")
	e := echo.New()
	e.POST("/getTemp", func(c echo.Context) error {
		req := new(Request)
		if err := c.Bind(req); err != nil {
			err := Error{"Ошибка парсинга body!"}
			return c.JSON(http.StatusOK, err)
		}
		lat, errLat := strconv.ParseFloat(req.Lat, 64)
		lon, errLon := strconv.ParseFloat(req.Lon, 64)

		if errLat == nil && errLon == nil {
			var weather, _ = weather.GetWeather(yandexWeatherApiKey, float32(lat), float32(lon))
			return c.JSON(
				http.StatusOK,
				InfoResponse{CurrTemp: fmt.Sprintf("%d", weather.Fact.Temp)})
		} else {
			err := Error{"Ошибка парсинга query lat/lan!"}
			return c.JSON(http.StatusOK, err)
		}
	})
	e.Logger.Fatal(e.Start(":1323"))

	w, _ := weather.GetWeather(yandexWeatherApiKey, 55.820501, 37.572370)
	fmt.Println(w)

}
