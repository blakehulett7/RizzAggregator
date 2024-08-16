package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    "localhost:" + os.Getenv("PORT"),
		Handler: mux,
	}
	fmt.Println("Christ is King!, also the server is starting...")
	mux.HandleFunc("GET /v1/healthz", ReportHealth)
	server.ListenAndServe()
}

func JsonResponse(writer http.ResponseWriter, statusCode int, responseData []byte) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(responseData)
}

func ReportHealth(writer http.ResponseWriter, request *http.Request) {

}
