package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/blakehulett7/RizzAggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Christ is King!, also the server is starting...")
	godotenv.Load()
	dbURL := os.Getenv("CONN")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		error := fmt.Errorf("Error opening database: %v", err)
		fmt.Println(error)
		return
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
	config.Database.NukeDB(context.Background())
	mux.HandleFunc("GET /v1/healthz", ReportHealth)
	mux.HandleFunc("POST /v1/users", config.AddUser)
	mux.HandleFunc("GET /v1/users", config.GetUser)
	server.ListenAndServe()
}

func JsonResponse(writer http.ResponseWriter, statusCode int, responseData []byte) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(responseData)
}
