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

	c := make(chan string)

	for _, link := range links {
		go makeRequest(link, c)
	}
	fmt.Println(<-c)
}

func makeRequest(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- "The link could be down"
		return
	}
	fmt.Println(link, "is up")
	c <- "The link is working"
}
