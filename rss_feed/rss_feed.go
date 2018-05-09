package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type Enclosure struct {
	Url    string `xml:"url,attr"`
	Length int64  `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type Item struct {
	Title     string    `xml:"title"`
	Link      string    `xml:"link"`
	Desc      string    `xml:"description"`
	Guid      string    `xml:"guid"`
	Enclosure Enclosure `xml:"enclosure"`
	PubDate   string    `xml:"pubDate"`
}

type Channel struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Desc  string `xml:"description"`
	Items []Item `xml:"item"`
}

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Config struct {
	Url        string `json:"url"`
	NumOfFeeds int    `json:"numOfFeeds"`
}

func main() {
	//get path of Executable
	ex, err := os.Executable()
	exPath := filepath.Dir(ex)
	//read config file and parse to struct
	configFilepath := exPath + "/rss_config.json"
	file, err := os.Open(configFilepath)
	if err != nil {
		fmt.Printf("Error open config file: %v\n", err)
		return
	}
	config := Config{}
	decoderJSON := json.NewDecoder(file)
	err = decoderJSON.Decode(&config)
	if err != nil {
		fmt.Printf("Error parse config to json: %v\n", err)
		return
	}

	//get rss feed from url
	response, err := http.Get(config.Url)
	if err != nil {
		fmt.Printf("Error http.GET(%v): %v\n", config.Url, err)
		return
	}
	defer response.Body.Close()

	//decode XML string from response to struct Rss
	rss := Rss{}
	decoderXML := xml.NewDecoder(response.Body)
	err = decoderXML.Decode(&rss)
	if err != nil {
		fmt.Printf("Error Decode XML: %v\n", err)
		return
	}

	//print top n feeds
	fmt.Println(rss.Channel.Title)
	fmt.Println("---\n")
	numOfFeeds := config.NumOfFeeds
	for i := 0; i < numOfFeeds; i++ {
		fmt.Println(i+1, ". ", rss.Channel.Items[i].Title, ": \n")
		fmt.Println(rss.Channel.Items[i].Desc, "\n")
		fmt.Println(rss.Channel.Items[i].Link, "\n")
		fmt.Println("---\n")
	}
}
