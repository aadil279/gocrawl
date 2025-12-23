package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var AllowedProtocols = [...]string{
	"http",
	"https",
}

func main() {
	/*reader := bufio.NewReader(os.Stdin)
	fmt.Print("URL to crawl: ")
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input")
	}*/

	CrawlSite("https://www.webscraper.io")
}

func CrawlSite(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error")
	}
	defer resp.Body.Close()

	tagList, _ := aTags(resp.Body)

	fmt.Println(tagList)
}

func aTags(body io.ReadCloser) ([]*html.Node, error) {
	doc, err := html.Parse(body)

	if err != nil {
		log.Fatal(err)
	}

	// List of scraped <a> tags
	var aTagList []*html.Node = []*html.Node{}

	// Loop through all the nodes in the document
	for n := range doc.Descendants() {

		// Check if the node is an <a> tag
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			// Append tag to list
			aTagList = append(aTagList, n)

			// Loop through the attributes of the <a> tag
			for _, a := range n.Attr {

				// Get the href attribute of the <a> tag
				if a.Key == "href" {
					val := a.Val
					protocol := strings.Split(val, `:`)[0]

					if protocolAllowed(protocol) {
						handleScrapedLink(string(a.Val))
					}
				}
			}
		}
	}

	return aTagList, nil
}

func handleScrapedLink(link string) {
	fmt.Println(link)
}

func protocolAllowed(proto string) bool {
	for _, p := range AllowedProtocols {
		if p == proto {
			return true
		}
	}
	return false
}
