package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/blakehulett7/RizzAggregator/internal/database"
	"github.com/google/uuid"
)

type apiConfig struct {
	Database *database.Queries
}

func ReportHealth(writer http.ResponseWriter, request *http.Request) {
	responseStruct := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}
	responseData, _ := json.Marshal(responseStruct)
	JsonResponse(writer, 200, responseData)
}

func (config apiConfig) AddUser(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Creating user..")
	decoder := json.NewDecoder(request.Body)
	clientParams := struct {
		Name string `json:"name"`
	}{}
	decoder.Decode(&clientParams)
	id := uuid.New()
	createdAt := time.Now()
	updatedAt := time.Now()
	name := clientParams.Name
	userStruct := database.CreateUserParams{
		ID:        id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Name:      name,
	}
	user, _ := config.Database.CreateUser(request.Context(), userStruct)
	responseData, _ := json.Marshal(user)
	JsonResponse(writer, 201, responseData)
}

func (config apiConfig) GetUser(writer http.ResponseWriter, request *http.Request) {
	apiToken := request.Header.Get("Authorization")
	apiKey, _ := strings.CutPrefix(apiToken, "ApiKey ")
	apiKey = strings.ReplaceAll(apiKey, "\"", "")
	user, err := config.Database.GetUser(request.Context(), apiKey)
	if err != nil {
		fmt.Println(err)
		JsonHeaderResponse(writer, 401)
		return
	}
	responseData, _ := json.Marshal(user)
	JsonResponse(writer, 200, responseData)
}

func Authenticator(config apiConfig, request *http.Request) (isAuthenticated bool, userID uuid.UUID) {
	apiToken := request.Header.Get("Authorization")
	apiKey, _ := strings.CutPrefix(apiToken, "ApiKey ")
	apiKey = strings.ReplaceAll(apiKey, "\"", "")
	user, err := config.Database.GetUser(request.Context(), apiKey)
	if err != nil {
		fmt.Println(err)
		return false, uuid.Nil
	}
	return true, user.ID
}

func (config apiConfig) AddFeed(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Calling feed creator...")
	isAuthenticated, userID := Authenticator(config, request)
	if !isAuthenticated {
		JsonHeaderResponse(writer, 401)
		return
	}
	decoder := json.NewDecoder(request.Body)
	clientParams := struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}{}
	decoder.Decode(&clientParams)
	userStruct := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      clientParams.Name,
		Url:       clientParams.Url,
		UserID:    userID,
	}
	feed, err := config.Database.CreateFeed(request.Context(), userStruct)
	if err != nil {
		fmt.Println(err)
		return
	}
	followStruct := database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
		FeedID:    feed.ID,
	}
	autoFollow, err := config.Database.CreateFeedFollows(request.Context(), followStruct)
	responseParams := struct {
		Feed       database.Feed       `json:"feed"`
		FeedFollow database.FeedFollow `json:"feed_follow"`
	}{
		feed,
		autoFollow,
	}
	responseData, _ := json.Marshal(responseParams)
	JsonResponse(writer, 201, responseData)
}

func (config apiConfig) GetFeeds(writer http.ResponseWriter, request *http.Request) {
	feedArray, err := config.Database.GetFeeds(request.Context())
	if err != nil {
		fmt.Println(err)
		return
	}
	responseData, _ := json.Marshal(feedArray)
	JsonResponse(writer, 200, responseData)
}

func (config apiConfig) AddFeedFollow(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("I have been called!")
	isAuthenticated, userID := Authenticator(config, request)
	if !isAuthenticated {
		JsonHeaderResponse(writer, 401)
		return
	}
	decoder := json.NewDecoder(request.Body)
	clientParams := struct {
		FeedID string `json:"feed_id"`
	}{}
	err := decoder.Decode(&clientParams)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(clientParams.FeedID)
	feedID, err := uuid.Parse(clientParams.FeedID)
	if err != nil {
		fmt.Println(err)
		fmt.Println("feed id not parsed", clientParams.FeedID)
		return
	}
	userStruct := database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
		FeedID:    feedID,
	}
	feedFollow, err := config.Database.CreateFeedFollows(request.Context(), userStruct)
	if err != nil {
		fmt.Println(err)
		fmt.Println("feed follow not created")
		return
	}
	responseData, _ := json.Marshal(feedFollow)
	JsonResponse(writer, 201, responseData)
}
