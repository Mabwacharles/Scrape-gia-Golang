package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func writeFile(data, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	file.WriteString(data)
}
func main() {
	url := "https://techcrunch.com/"

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

	success := doc.Find("div.river").Find("div.post-block").Each(func(index int, item *goquery.Selection) {
		h2 := item.Find("h2").Text() // get the title
		p := item.Find("p").Text()   // get the description
		fmt.Println(h2)
		fmt.Println(p)
	})
	// fmt.Println(doc)
	// fmt.Println(success)

	fmt.Println("Successfully retrieved", success)

	// writeFile(success, "techcrunch.html")
}
