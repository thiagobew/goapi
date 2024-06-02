package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/thiagobew/goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(r chi.Router) {
		r.Use(middleware.Authorization)
		r.Get("/coins", GetCoinBalance)
	})
}