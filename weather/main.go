package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	api_key := "5facba76f0064528bb4162342242810"
	request, err := http.Get("http://api.weatherapi.com/v1/current.json?key=" + api_key + "&q=London")
	if err != nil {
		return
	}
	defer request.Body.Close()
	newFile, err2 := os.Create("smth.json")
	if err2 != nil {
		return
	}
	defer newFile.Close()
	_, err3 := io.Copy(newFile, request.Body)
	if err3 != nil {
		return
	}

}