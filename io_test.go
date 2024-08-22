package main

import (
	"fmt"
	"testing"
)

func TestRizzFetcher(t *testing.T) {
	FetchFeed("https://www.dailywire.com/feeds/rss.xml")
	FetchFeed("https://www.mtggoldfish.com/feed")
}

func TestRizzProcessor(t *testing.T) {
	rizzArray := []Rss{
		FetchFeed("https://blog.boot.dev/index.xml"),
		FetchFeed("https://wagslane.dev/index.xml"),
		FetchFeed("https://www.dailywire.com/feeds/rss.xml"),
		FetchFeed("https://www.mtggoldfish.com/feed"),
	}
	for idx, rizz := range rizzArray {
		fmt.Println("\nShowing rizz", idx)
		ProcessRizz(rizz)
		fmt.Println()
	}
}
