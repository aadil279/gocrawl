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

	aTags(resp.Body)
}

func aTags(body io.ReadCloser) {
	doc, err := html.Parse(body)

	if err != nil {
		log.Fatal(err)
	}

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					val := a.Val
					protocol := strings.Split(val, `:`)[0]

					if protocolAllowed(protocol) {
						fmt.Println(string(a.Val))
					}
				}
			}
		}
	}
}

func protocolAllowed(proto string) bool {
	for _, p := range AllowedProtocols {
		if p == proto {
			return true
		}
	}
	return false
}
