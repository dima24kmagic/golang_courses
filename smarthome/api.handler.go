package smarthome

import "net/http"

//GlobalAPI is @string api path, that using for fetching forecast
var GlobalAPI = "https://api.darksky.net/forecast/c1a755c1be234b586aa61773878dfbfd/53.9,-27.57"

//GetWeather - fetch weather from api
func GetWeather() string {
	http.Get(GlobalAPI)
	return "Weather is fetched"
}
