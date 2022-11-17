package main

import (
	fattyauth "fatty/internal/fatty-auth"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func getPrivateRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return r
}

func main() {
	godotenv.Load()

	authService := fattyauth.InitAuth()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	m := authService.Middleware()
	r.With(m.Auth).Mount("/api", getPrivateRoutes())

	authRoutes, avatarRoutes := authService.Handlers()
	r.Mount("/auth", authRoutes)
	r.Mount("/avatar", avatarRoutes)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", r)
}