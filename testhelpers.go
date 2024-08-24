package main

import (
	"context"
	"fmt"
	"time"

	"github.com/blakehulett7/RizzAggregator/internal/database"
	"github.com/google/uuid"
)

func CreateSampleUsers(config apiConfig) []database.User {
	Blake, err := config.Database.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Blake",
	})
	if err != nil {
		fmt.Println("Couldn't create user Blake")
		return []database.User{}
	}
	Brett, err := config.Database.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Brett",
	})
	if err != nil {
		fmt.Println("Couldn't create user Brett")
		return []database.User{}
	}
	Bo, err := config.Database.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Bo",
	})
	if err != nil {
		fmt.Println("Couldn't create user Bo")
		return []database.User{}
	}
	return []database.User{Blake, Brett, Bo}
}

func CreateSampleFeeds(config apiConfig, user1, user2, user3 database.User) []database.Feed {
	feed1, err := config.Database.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Feed 1",
		Url:       "Url1.com",
		UserID:    user1.ID,
	})
	if err != nil {
		fmt.Println(err)
	}
	feed2, err := config.Database.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Feed 2",
		Url:       "Url2.com",
		UserID:    user2.ID,
	})
	if err != nil {
		fmt.Println(err)
	}
	feed3, err := config.Database.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Feed 3",
		Url:       "Url3.com",
		UserID:    user3.ID,
	})
	if err != nil {
		fmt.Println(err)
	}
	return []database.Feed{feed1, feed2, feed3}
}

func CreateSampleFollows(config apiConfig, user1, user2, user3 database.User, feed1, feed2, feed3 database.Feed) []database.FeedFollow {
	follow1, err := config.Database.CreateFeedFollows(context.Background(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user1.ID,
		FeedID:    feed1.ID,
	})
	if err != nil {
		fmt.Println(err)
	}
	follow2, err := config.Database.CreateFeedFollows(context.Background(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user1.ID,
		FeedID:    feed2.ID,
	})
	if err != nil {
		fmt.Println(err)
	}
	follow3, err := config.Database.CreateFeedFollows(context.Background(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user1.ID,
		FeedID:    feed3.ID,
	})
	if err != nil {
		fmt.Println(err)
	}
	follow4, err := config.Database.CreateFeedFollows(context.Background(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user2.ID,
		FeedID:    feed1.ID,
	})
	if err != nil {
		fmt.Println(err)
	}
	follow5, err := config.Database.CreateFeedFollows(context.Background(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user2.ID,
		FeedID:    feed2.ID,
	})
	if err != nil {
		fmt.Println(err)
	}
	follow6, err := config.Database.CreateFeedFollows(context.Background(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user3.ID,
		FeedID:    feed3.ID,
	})
	if err != nil {
		fmt.Println(err)
	}
	return []database.FeedFollow{follow1, follow2, follow3, follow4, follow5, follow6}
}
