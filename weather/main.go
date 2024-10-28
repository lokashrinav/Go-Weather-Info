package main

import (
	"fmt"
	"io"
	"net/http"
	//"os"
	"encoding/json"
)

type WeatherData struct {
    Location struct {
        Name string `json:"name"` 
    } `json:"location"`
    Current struct {
        TempF    float64 `json:"temp_f"`
        Humidity int     `json:"humidity"`
        Condition struct {
            Text string `json:"text"`
        } `json:"condition"` 
    } `json:"current"`
}

func main() {
	api_key := "5facba76f0064528bb4162342242810"
	request, err := http.Get("http://api.weatherapi.com/v1/current.json?key=" + api_key + "&q=London")
	if err != nil {
		return
	}
	defer request.Body.Close()
	/*newFile, err2 := os.Create("smth.json")
	if err2 != nil {
		return
	}
	defer newFile.Close()
	_, err3 := io.Copy(newFile, request.Body)
	if err3 != nil {
		return
	}*/
	body, err := io.ReadAll(request.Body)

	var x WeatherData
    err2 := json.Unmarshal([]byte(body), &x)
	if err2 != nil {
		return
	}
	
	fmt.Println("Location:", x.Location.Name)
	fmt.Printf("Temperature: %.2f°F\n", x.Current.TempF)
	fmt.Println("Humidity:", x.Current.Humidity)
	fmt.Println("Condition:", x.Current.Condition.Text)
}