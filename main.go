package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/blakehulett7/RizzAggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("CONN")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Errorf("Error opening database", err)
	}
	dbQueries := database.New(db)
	config := apiConfig{
		dbQueries,
	}
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    "localhost:" + os.Getenv("PORT"),
		Handler: mux,
	}
	fmt.Println("Christ is King!, also the server is starting...")
	mux.HandleFunc("GET /v1/healthz", ReportHealth)
	mux.HandleFunc("POST /v1/healthz", config.AddUser)
	server.ListenAndServe()
}

func JsonResponse(writer http.ResponseWriter, statusCode int, responseData []byte) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(responseData)
}
