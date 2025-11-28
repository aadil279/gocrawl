package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.webscraper.io")

	if err != nil {
		fmt.Println("Error")
	}
	defer resp.Body.Close()

	rb, err := io.ReadAll(resp.Body)

	fmt.Println(string(rb))
}
