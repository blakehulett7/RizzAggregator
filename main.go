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
	server.ListenAndServe()
}
