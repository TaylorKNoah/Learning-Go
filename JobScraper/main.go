package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

//1) get pages
//2) visit pages
//3) Extract info

var baseURL string = "https://www.indeed.com/jobs?q=machine+learning&limit=50"

func main() {
	pages := getPages()

	for i := 0; i < pages; i++ {
		getPage(i)
	}
}

//checks error status
func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//checks http response code
func checkStatusCode(resp *http.Response) {
	if resp.StatusCode != 200 {
		log.Fatalln("Request failed with status: ", resp.StatusCode)
	}
}

//finds the number of urls in a page
// eg: 2, 3, 4, 5, nextpg->
func getPages() int {

	pages := 0

	resp, err := http.Get(baseURL)
	checkError(err)
	checkStatusCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkError(err)

	doc.Find(".pagination")
	doc.Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages

}

func getPage(page int) {
	pageURL := baseURL + "start=" + strconv.Itoa(page*50)

	resp, err := http.Get(pageURL)
	checkError(err)
	checkStatusCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkError(err)

	id := ""
	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, s *goquery.Selection) {
		id, _ = s.Attr("data-jk")
	})

	fmt.Println(id)
}
