package main

import (
	"log"
	gohttp "net/http"
	_ "over-engineered-calculator/docs"
	"over-engineered-calculator/internal"
	"over-engineered-calculator/internal/helpers"
)

// SetupHttp initializes the HTTP router and sets up the routes for the API.
// @title Over Engineered Calculator
// @version 1.0
// @description This is a sample API for an over engineered calculator for a coding test
// @BasePath /api/v1
func main() {
	r := internal.SetupHttp()
	port := helpers.MustFromEnv("PORT", "8080")

	log.Printf("🚀 Server running on port %s", port)
	log.Fatal(gohttp.ListenAndServe(":"+port, r))
}
