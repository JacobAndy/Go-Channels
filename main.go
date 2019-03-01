package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// creating a new channel
	c := make(chan string)

	for _, link := range links {
		// using 'go' creates a new thread aka go routine
		// this is a child go routine
		// only use go key words in front of function calls
		go checkLink(link, c)
	}

	//infinite loop, never going to exit
	// for{}

	// wait for a value to come to the channel, assign it to l and then run the body of the loop
	for l := range c {
		// we are waiting for a message to come through the channel
		// before we go to the next iteration loop
		go func(link string) {
			time.Sleep(time.Second * 3)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
