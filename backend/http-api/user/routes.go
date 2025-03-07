package http_api_user

import (
	"net/http"

	"github.com/devflex-pro/tma-starter-kit/backend/domain"
	"github.com/devflex-pro/tma-starter-kit/backend/http-api/middlewares"
)

func RegisterRoutes(
	mux *http.ServeMux,
	db domain.UserDB,
	botToken string,
) {

	h := NewUserHandler(db)
	mux.HandleFunc(
		"/api/users",
		middlewares.TelegramAuthMiddleware(botToken, h.Save),
	)
	mux.HandleFunc(
		"/api/users/{id}",
		middlewares.TelegramAuthMiddleware(botToken, h.Read),
	)
}
