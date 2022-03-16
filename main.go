package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://gamesindustryafrica.com/"

	// Create a new request using http
	response, error := http.Get(url)

	if error != nil {
		fmt.Println(error)
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		fmt.Println("Successfully retrieved", url)
	} else {
		fmt.Println("Couldn't retrieve", url)
	}

	doc, error := goquery.NewDocumentFromReader(response.Body)
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(doc)
}
