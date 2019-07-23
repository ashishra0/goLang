package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.youtube.com",
		"http://www.golang.org",
		"http://www.gmail.com",
	}

	for _, link := range links {
		makeRequest(link)
	}
}

func makeRequest(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		return
	}
	fmt.Println(link, "is up")
}
