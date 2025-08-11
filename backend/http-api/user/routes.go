package http_api_user

import (
    "net/http"
    "os"

    "github.com/devflex-pro/tma-starter-kit/backend/domain"
    "github.com/devflex-pro/tma-starter-kit/backend/http-api/middlewares"
)

func RegisterRoutes(
	mux *http.ServeMux,
	db domain.UserDB,
	botToken string,
) {

    h := NewUserHandler(db)
    if os.Getenv("BYPASS_TELEGRAM_AUTH") == "true" {
        mux.HandleFunc("/api/users", h.Save)
        // Note: trailing slash pattern to match subpaths like /api/users/123
        mux.HandleFunc("/api/users/", h.Read)
    } else {
        mux.HandleFunc(
            "/api/users",
            middlewares.TelegramAuthMiddleware(botToken, h.Save),
        )
        mux.HandleFunc(
            "/api/users/",
            middlewares.TelegramAuthMiddleware(botToken, h.Read),
        )
    }
}
