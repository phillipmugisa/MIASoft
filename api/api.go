package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/cors"
)

func (a *AppServer) Run() *http.Server {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	a.registerRoutes(r)

	return &http.Server{
		Addr:           fmt.Sprintf(":%s", a.port),
		Handler:        r,
		ReadTimeout:    10 & time.Second,
		WriteTimeout:   10 & time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func (a *AppServer) registerRoutes(r *chi.Mux) {
	// r.Post("/", MakeApiHandler(a.HomeHandler))
	r.Post("/", MakeApiHandler(a.WelcomeHandler))
	// r.Post("/auth", MakeApiHandler(a.AuthHandler))
	// r.Post("/init", MakeApiHandler(a.SubcriptionInitializationHandler))
	// r.Post("/init/companies", MakeApiHandler(a.HomeHandler))
}
