package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	healthHandler := handler.NewHealthzHandler()
	todoHandler := handler.NewTODOHandler(service.NewTODOService(todoDB))
	mux.HandleFunc("/healthz", healthHandler.ServeHTTP)
	mux.HandleFunc("/todos", todoHandler.ServeHTTP)

	return mux
}
