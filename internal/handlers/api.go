package handlers

import (
    "github.com/go-chi/chi"
    chimiddle "github.com/go-chi/chi/middleware"
    "github.com/jjetta/go-api-tutorial/internal/middleware"
)

func Handler(r *chi.Mux) {
    // Global middleware (means that this middleware is applied to all enpoints made)
    r.Use(chimiddle.StripSlashes) // (StripSlashes ignores trailing slahes on an endpoint) 
    r.Route("/account", func(router chi.Router) {

        // Middleware for /account route
        router.Use(middleware.Authorization)

        router.Get("/coins", GetCoinBalance) 
    })
}


