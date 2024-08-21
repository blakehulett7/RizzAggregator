package main

import (
	"database/sql"
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
		ID:            uuid.New(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Name:          clientParams.Name,
		Url:           clientParams.Url,
		UserID:        userID,
		LastFetchedAt: sql.NullTime{},
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
		JsonHeaderResponse(writer, 400)
		return
	}
	feedID, err := uuid.Parse(clientParams.FeedID)
	if err != nil {
		fmt.Println(err)
		JsonHeaderResponse(writer, 400)
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
		JsonHeaderResponse(writer, 400)
		return
	}
	responseData, _ := json.Marshal(feedFollow)
	JsonResponse(writer, 201, responseData)
}

func (config apiConfig) DeleteFeedFollow(writer http.ResponseWriter, request *http.Request) {
	isAuthenticated, userID := Authenticator(config, request)
	if !isAuthenticated {
		JsonHeaderResponse(writer, 401)
		return
	}
	feedFollowString := request.PathValue("feedFollowID")
	feedFollowString = strings.ReplaceAll(feedFollowString, "\"", "")
	feedFollowID, err := uuid.Parse(feedFollowString)
	if err != nil {
		fmt.Println(err)
		return
	}
	params := database.DeleteFollowParams{
		ID:     feedFollowID,
		UserID: userID,
	}
	err = config.Database.DeleteFollow(request.Context(), params)
	if err != nil {
		fmt.Println(err)
		JsonHeaderResponse(writer, 400)
	}
	JsonHeaderResponse(writer, 200)
}

func (config apiConfig) GetFeedFollows(writer http.ResponseWriter, request *http.Request) {
	isAuthenticated, userID := Authenticator(config, request)
	if !isAuthenticated {
		JsonHeaderResponse(writer, 401)
		return
	}
	followsArray, err := config.Database.GetFollows(request.Context(), userID)
	if err != nil {
		fmt.Println(err)
		JsonHeaderResponse(writer, 400)
		return
	}
	responseData, err := json.Marshal(followsArray)
	if err != nil {
		fmt.Println(err)
		JsonHeaderResponse(writer, 400)
		return
	}
	JsonResponse(writer, 200, responseData)
}

func (config apiConfig) RunTests(writer http.ResponseWriter, request *http.Request) {
	fetchQueue, err := config.Database.GetNextFeedsToFetch(request.Context(), 2)
	if err != nil {
		fmt.Println("couldn't update fetch queue...")
		JsonHeaderResponse(writer, 400)
		return
	}
	responseData, err := json.Marshal(fetchQueue)
	if err != nil {
		fmt.Println("couldn't marshal fetchArray...")
		JsonHeaderResponse(writer, 400)
		return
	}
	JsonResponse(writer, 200, responseData)
}
