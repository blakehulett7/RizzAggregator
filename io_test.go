package main

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/blakehulett7/RizzAggregator/internal/database"
	"github.com/google/uuid"
)

func TestRizzFetcher(t *testing.T) {
	FetchFeed("https://www.dailywire.com/feeds/rss.xml")
	FetchFeed("https://www.mtggoldfish.com/feed")
}

func TestRizzProcessor(t *testing.T) {
	config := OpenDB()
	defer config.Database.NukeUsersDB(context.Background())
	defer config.Database.NukeFeedsDB(context.Background())
	defer config.Database.NukePostsDB(context.Background())
	user, err := config.Database.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Blake",
	})
	if err != nil {
		t.Fatal(err)
	}
	targetUrls := []string{"https://blog.boot.dev/index.xml",
		"https://wagslane.dev/index.xml",
		"https://www.dailywire.com/feeds/rss.xml",
		"https://www.mtggoldfish.com/feed"}
	feedArray := []database.Feed{}
	for idx, url := range targetUrls {
		feed, err := config.Database.CreateFeed(context.Background(), database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      strconv.Itoa(idx),
			Url:       url,
			UserID:    user.ID,
		})
		if err != nil {
			t.Fatal(err)
		}
		feedArray = append(feedArray, feed)
	}
	for _, feed := range feedArray {
		fmt.Println("Fetching rizz from:", feed.Url)
		rizz := FetchFeed(feed.Url)
		fmt.Println("Processing the rizz...")
		ProcessRizz(feed, config, rizz)
		fmt.Println()
	}
}
