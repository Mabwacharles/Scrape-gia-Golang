package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

// Create a new request using http
func getHtml(url string) *http.Response {

	response, error := http.Get(url)

	if error != nil {
		fmt.Println(error)
	}

	if response.StatusCode == 200 {
		fmt.Println("Successfully retrieved", url)
	} else {
		fmt.Println("Couldn't retrieve", url)
	}

	return response

}

func writeCsv(posts []string) {
	filename := "posts.csv"

	file, error := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if error != nil {
		fmt.Println(error)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	error = writer.Write(posts)
	if error != nil {
		fmt.Println(error)
	}
}

func scrapePageData(doc *goquery.Document) {
	doc.Find("div.river").Find("div.post-block").Each(func(index int, item *goquery.Selection) {
		h2 := item.Find("h2").Text() // get the title
		p := item.Find("p").Text()   // get the description
		url := item.Find("a").AttrOr("href", "")

		excerpt := strings.TrimSpace(item.Find("div.post-block_content").Text())

		posts := []string{h2, p, excerpt, url}

		writeCsv(posts)

	})
}

func main() {

	url := "https://techcrunch.com/"

	response := getHtml(url)
	defer response.Body.Close()

	doc, error := goquery.NewDocumentFromReader(response.Body)
	if error != nil {
		fmt.Println(error)

		scrapePageData(doc)
	}

}
