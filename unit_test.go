package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/blakehulett7/RizzAggregator/internal/database"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func OpenDB() apiConfig {
	fmt.Println("Christ is King!, also the test is starting...")
	godotenv.Load()
	dbURL := os.Getenv("CONN")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		error := fmt.Errorf("Error opening database: %v", err)
		fmt.Println(error)
		return apiConfig{}
	}
	dbQueries := database.New(db)
	config := apiConfig{
		dbQueries,
	}
	return config
}

func (config apiConfig) CreateSampleFeeds() []database.Feed {
	userID1 := uuid.New()
	_, err := config.Database.CreateUser(context.Background(), database.CreateUserParams{
		ID:        userID1,
		CreatedAt: time.Date(1997, time.January, 25, 0, 0, 0, 0, time.FixedZone("central", 0)),
		UpdatedAt: time.Now(),
		Name:      "Blake",
	})
	if err != nil {
		fmt.Println(err)
	}
	feed1, err := config.Database.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:            uuid.New(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Name:          "The Boot.dev Blog",
		Url:           "https://blog.boot.dev/index.xml",
		UserID:        userID1,
		LastFetchedAt: sql.NullTime{},
	})
	if err != nil {
		fmt.Println(err)
	}
	feed2, err := config.Database.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Wagslane Blog",
		Url:       "https://wagslane.dev/index.xml",
		UserID:    userID1,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	feed3, err := config.Database.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "DailyWire",
		Url:       "https://www.dailywire.com/feeds/rss.xml",
		UserID:    userID1,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	feed4, err := config.Database.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "MTG Goldfish",
		Url:       "https://www.mtggoldfish.com/feed",
		UserID:    userID1,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	return []database.Feed{feed1, feed2, feed3, feed4}
}

func TestUpdateFetchQueue(t *testing.T) {
	config := OpenDB()
	defer config.Database.NukeUsersDB(context.Background())
	defer config.Database.NukeFeedsDB(context.Background())
	feedArray := config.CreateSampleFeeds()
	result1 := []database.Feed{feedArray[0], feedArray[1]}
	result2 := []database.Feed{feedArray[0], feedArray[1], feedArray[2]}
	fetchQueue1, err := config.Database.GetNextFeedsToFetch(context.Background(), 2)
	if !reflect.DeepEqual(fetchQueue1, result1) || err != nil {
		t.Log("Err:", err)
		t.Fatal("Failed to update fetch queue")
	}
	fetchQueue2, err := config.Database.GetNextFeedsToFetch(context.Background(), 3)
	if !reflect.DeepEqual(fetchQueue2, result2) || err != nil {
		t.Log("Err:", err)
		t.Fatal("Failed to update fetch queue 2")
	}
	fmt.Println("Successfully updated fetch queues!")
}
