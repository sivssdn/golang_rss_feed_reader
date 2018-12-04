package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"fmt"
)

/*For Medium Rss Feed*/
type MediumRoot struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel
}

type Channel struct{
	XMLName		xml.Name `xml:"channel"`
	Item []Item `xml:"item"`
}

type Item struct {
	XMLName       xml.Name `xml:"item"`
	Title         string   `xml:"title"`
	Link		  string    `xml:"link"`
	Category	  []string   `xml:"category"`
	PublishedDate  string    `xml:"pubDate"`
	/*Content		  string   `xml:"encoded"`*/
}


func main() {
	crawlMediumFeeds()	
	
}

func crawlMediumFeeds(){
	/*list the links of all feed urls in the rss_feed_urls array */
	var rss_feed_urls = []string{""}

	for _, rss_feed_url := range rss_feed_urls{
		
		client := &http.Client{}
		request,_ := http.NewRequest("GET", rss_feed_url, nil)
		request.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.92 Safari/537.36")
		response, _ := client.Do(request)
		defer response.Body.Close()
		response_body, _ := ioutil.ReadAll(response.Body)

		var mediumRss MediumRoot
		xml.Unmarshal(response_body, &mediumRss)

		for i := 0; i < len(mediumRss.Channel.Item); i++ {
			fmt.Println("" + mediumRss.Channel.Item[i].Link)
		}
	}


}
