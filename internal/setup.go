package internal

import (
	"github.com/go-chi/chi/v5"
	http_middleware "github.com/sfrand8/over-engineered-calculator/pkg/http-middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"over-engineered-calculator/internal/features/calculate"
	"over-engineered-calculator/internal/features/get_history"
)

func SetupHttp() *chi.Mux {
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
