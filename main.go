package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://gamesindustryafrica.com/"

	// Create a new request using http
	response, error := http.Get(url)
	if error != nil {
		fmt.Println(error)
	}
}
