package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func checkURLStatus(url string, c chan string) {
	_, err := http.Get(url)
	time.Sleep(time.Second)
	if err != nil {
		fmt.Println(url, "might be down!")
		c <- url
		return
	}

	fmt.Println(url, "is up!")
	c <- url
}

func main() {
	links := os.Args[1:]
	c := make(chan string)
	if len(links) == 0 {
		fmt.Println("Usage: ./main http://www.example.com http://www.example.org")
		os.Exit(1)
	}
	for _, link := range links {
		go checkURLStatus(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second)
			go checkURLStatus(link, c)
		}(l)
	}
}
