package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func InitializeServerRoutes(srv *Server) http.Handler {
	router := mux.NewRouter()

	handlersCors := cors.New(cors.Options{
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "pragma", "X-Organization"},
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Total-Count"},
		MaxAge:           300,
	})

	// Authentication routes
	auth := router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", srv.HandleLogin).Methods(http.MethodPost)
	auth.HandleFunc("/register", srv.HandleRegister).Methods(http.MethodPost)

	// Api v1 routes
	_ = router.PathPrefix("/api/v1").Subrouter()

	appRouter := handlersCors.Handler(router)
	return appRouter
}
