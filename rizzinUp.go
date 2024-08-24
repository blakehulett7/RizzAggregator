package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/blakehulett7/RizzAggregator/internal/database"
	"github.com/google/uuid"
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

func ProcessRizz(feed database.Feed, config apiConfig, rizzStruct Rss) {
	for _, post := range rizzStruct.Channel.Item {
		published, err := time.Parse(time.RFC1123Z, post.PubDate)
		if err != nil {
			published, err = time.Parse(time.RFC1123, post.PubDate)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
		config.Database.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     post.Title,
			Url:       post.Link,
			Description: sql.NullString{
				String: post.Description,
				Valid:  true,
			},
			PublishedAt: published,
			FeedID:      feed.ID,
		})
	}
	for _, post := range rizzStruct.Entry {
		published, err := time.Parse(time.RFC3339, post.Published)
		if err != nil {
			fmt.Println(err)
			continue
		}
		config.Database.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     post.Title,
			Url:       post.URL,
			Description: sql.NullString{
				String: post.Summary,
				Valid:  true,
			},
			PublishedAt: published,
			FeedID:      feed.ID,
		})
	}
}

func (config apiConfig) WorkTheRizz() {
	for {
		var waitGroup sync.WaitGroup
		fetchesAtOnce := 3
		fmt.Printf("Adding next %v feeds to queue...\n", fetchesAtOnce)
		fetchQueue, err := config.Database.GetNextFeedsToFetch(context.Background(), int32(fetchesAtOnce))
		waitGroup.Add(len(fetchQueue))
		if err != nil {
			fmt.Println(err)
		}
		for _, feed := range fetchQueue {
			go UpdateFeed(feed, config, &waitGroup)
		}
		waitGroup.Wait()
		time.Sleep(time.Second * 60)
	}
}

func UpdateFeed(feed database.Feed, config apiConfig, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Println("Fetching rizz from:", feed.Url)
	fmt.Println("")
	rss := FetchFeed(feed.Url)
	fmt.Println("Processing rizz returned from:", feed.Url)
	fmt.Println("")
	ProcessRizz(feed, config, rss)
	fmt.Println("")
	fmt.Println("Updating last fetched at and updated at for", feed.Url, "feed...")
	fmt.Println("")
	config.Database.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:        feed.ID,
		UpdatedAt: time.Now(),
	})
}
