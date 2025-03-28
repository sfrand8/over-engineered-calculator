package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sfrand8/over-engineered-calculator/pkg/http-middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	gohttp "net/http"
	_ "over-engineered-calculator/docs"
	"over-engineered-calculator/internal/features/calculate"
	"over-engineered-calculator/internal/features/get_history"
	"over-engineered-calculator/internal/helpers"
)

func main() {
	r := setupHttp()
	port := helpers.MustFromEnv("PORT", "8080")

	log.Printf("🚀 Server running on port %s", port)
	log.Fatal(gohttp.ListenAndServe(":"+port, r))
}

func setupHttp() *chi.Mux {
	httpMux := chi.NewRouter()
	httpMux.Route("/swagger", func(r chi.Router) {
		r.Get("/*", httpSwagger.WrapHandler)
	})

	httpMux.Route("/api/v1", func(r chi.Router) {
		r.Use(http_middleware.AuthMiddleware)
		calculate.Setup(r)
		get_history.Setup(r)
	})
	return httpMux
}
