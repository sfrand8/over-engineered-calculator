package main

import (
	"OverEngineeredCalculator/features/calculate"
	"OverEngineeredCalculator/features/get_history"
	"OverEngineeredCalculator/helpers"
	"OverEngineeredCalculator/http"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"

	_ "OverEngineeredCalculator/docs"
	gohttp "net/http"
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
		r.Use(http.AuthMiddleware)
		calculate.Setup(r)
		get_history.Setup(r)
	})
	return httpMux
}
