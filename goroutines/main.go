package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//fmt.Println("Hello, sailor!")
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://golang.org",
		"http://reddit.com",
		"http://fatter.org",
	}

	c := make(chan string)

	// Initial naive version
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// Infinite loop version
	// for {
	// 	go checkLink(<-c, c)
	// }

	// More canonical version
	// for l range c {
	// 	go checkLink(l, c)
	// }

	for _, link := range links {
		go checkLink(link, c)
	}

	// Pausing using function literal (a lambda)
	for l := range c {
		// ALWAYS pass by value, never share variables between coroutines!
		// Here we pass 'l' into our function literal
		go func(link string) {
			time.Sleep(5 * time.Second)
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
