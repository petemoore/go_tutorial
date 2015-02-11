package main

import (
	"fmt"
)

type comm struct {
	url   string
	reply chan bool
}

var (
	fetched map[string]bool = make(map[string]bool)
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan comm) {
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		reply := make(chan bool)
		ch <- comm{url: u, reply: reply}
		decision := <-reply
		if decision {
			fmt.Printf("Downloading from url %v\n", u)
			Crawl(u, depth-1, fetcher, ch)
		} else {
			fmt.Printf("Not downloading from url %v\n", u)
		}
	}
	return
}

func main() {
	// channel to send urls to
	ch := make(chan comm)
	go boss(ch)
	fetched["http://golang.org/"] = true
	Crawl("http://golang.org/", 4, fetcher, ch)
}

func boss(ch chan comm) {
	for {
		comm := <-ch
		if !fetched[comm.url] {
			fetched[comm.url] = true
			comm.reply <- true
		} else {
			comm.reply <- false
		}
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
