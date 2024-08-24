package main

import (
	"context"
	"database/sql"
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

func CreateSamplePosts(config apiConfig, numPosts int, feeds ...database.Feed) []database.Post {
	postArray := []database.Post{}
	for idx, feed := range feeds {
		for i := 1; i <= numPosts; i++ {
			title := fmt.Sprintf("Title %v", i)
			post, err := config.Database.CreatePost(context.Background(), database.CreatePostParams{
				ID:        uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Title:     title,
				Url:       fmt.Sprintf("Url %v, %v", idx, i),
				Description: sql.NullString{
					String: fmt.Sprintf("Description %v", 1),
					Valid:  true,
				},
				PublishedAt: time.Now(),
				FeedID:      feed.ID,
			})
			if err != nil {
				fmt.Printf("Error: %v happened on feed index: %v post number: %v", err, idx, i)
			}
			postArray = append(postArray, post)
		}
	}
	return postArray
}
