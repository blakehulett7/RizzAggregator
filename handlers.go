package main

import (
	"encoding/json"
	"net/http"

	"github.com/blakehulett7/RizzAggregator/internal/database"
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

func (config apiConfig) AddUser() {

}
