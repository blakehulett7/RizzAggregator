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
		return
	}
	responseData, _ := json.Marshal(user)
	JsonResponse(writer, 200, responseData)
}

func IsAuthenticated(config apiConfig, request *http.Request) (isAuthorized bool, userID string) {
	apiToken := request.Header.Get("Authorization")
	apiKey, _ := strings.CutPrefix(apiToken, "ApiKey ")
	apiKey = strings.ReplaceAll(apiKey, "\"", "")
	user, err := config.Database.GetUser(request.Context(), apiKey)
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	return true, user.ID.String()
}
