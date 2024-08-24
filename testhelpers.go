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
		ID: uuid.New(),
	})
	if err != nil {

	}
	return []database.Feed{feed1}
}
