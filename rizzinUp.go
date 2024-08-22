package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/blakehulett7/RizzAggregator/internal/database"
)

func FetchFeed(url string) Rss {
	responseData, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return Rss{}
	}
	body, err := io.ReadAll(responseData.Body)
	if err != nil {
		fmt.Println(err)
		return Rss{}
	}
	rss := Rss{}
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		fmt.Println(err)
		return Rss{}
	}
	return rss
}

func ProcessRizz(rizzStruct Rss) {
	titleArray := []string{}
	for _, post := range rizzStruct.Channel.Item {
		titleArray = append(titleArray, post.Title)
	}
	for _, post := range rizzStruct.Entry {
		titleArray = append(titleArray, post.Title)
	}
	for _, post := range titleArray {
		fmt.Println(post)
	}
}

func (config apiConfig) WorkTheRizz() {
	var waitGroup sync.WaitGroup
	fetchesAtOnce := 3
	waitGroup.Add(fetchesAtOnce)
	fmt.Printf("Adding next %v feeds to queue...\n", fetchesAtOnce)
	fetchQueue, err := config.Database.GetNextFeedsToFetch(context.Background(), int32(fetchesAtOnce))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, feed := range fetchQueue {
		go UpdateFeed(feed, config, &waitGroup)
	}
	waitGroup.Wait()
}

func UpdateFeed(feed database.Feed, config apiConfig, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Println("Fetching rizz from:", feed.Url)
	fmt.Println("")
	rss := FetchFeed(feed.Url)
	fmt.Println("Processing rizz returned from:", feed.Url)
	fmt.Println("")
	ProcessRizz(rss)
	fmt.Println("")
	fmt.Println("Updating last fetched at and updated at for", feed.Url, "feed...")
	fmt.Println("")
	config.Database.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:        feed.ID,
		UpdatedAt: time.Now(),
	})
}
