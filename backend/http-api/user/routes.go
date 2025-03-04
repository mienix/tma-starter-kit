package http_api_user

import (
	"net/http"

	"github.com/devflex-pro/tma-starter-kit/backend/domain"
)

func RegisterRoutes(
	mux *http.ServeMux,
	db domain.UserDB,
) {
	h := NewUserHandler(db)
	mux.HandleFunc("/api/users", h.Save)
	mux.HandleFunc("/api/users/{id}", h.Read)
}
