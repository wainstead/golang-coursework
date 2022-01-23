package firstmain

import (
	"fmt"
	"net/http"
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

	for _, link := range links {
		// fmt.Println(link
		go checkLink(link, c)
		// like forking, if you please.
		// The Go scheduler will only use one core by default.
		// "Concurrency is not parallelism" is a common refrain in Go-land.
		// We only get parallelism when we use two or more cores.
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		//fmt.Println(link, "might be down!")
		c <- link + " might be down!"
		return
	}
	// fmt.Println(link, "is up!")
	c <- link + " is up!"
}
