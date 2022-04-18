package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type News struct {
	Headline     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

func main() {
	var r RSS

	data := readGoogleTrends(&r)

	err := xml.Unmarshal(data, &r)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("GOOGLE DAILY TRENDS FOR TURKEY")
	fmt.Println("------------------------------")

	for i, v := range r.Channel.ItemList {
		fmt.Println("#", i+1, v.Title)
		fmt.Println(v.Link)
		fmt.Println(v.NewsItems[0].Headline)
		fmt.Println("\n------------------------------")
	}

}

func readGoogleTrends(r *RSS) []byte {
	response := getGoogleTrends()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	return data
}

func getGoogleTrends() *http.Response {
	response, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=TR")

	if err != nil {
		panic(err)
	}

	return response
}
