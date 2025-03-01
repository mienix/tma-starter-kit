package http_api_user

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/devflex-pro/tma-starter-kit/backend/domain"
)

type UserHandler struct {
	db domain.UserDB
}

func NewUserHandler(
	db domain.UserDB,
) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) Read(
	w http.ResponseWriter,
	r *http.Request,
) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[2] == "" {
		http.Error(
			w,
			"missing user ID",
			http.StatusBadRequest)
		return
	}
	id := parts[2]

	user, err := h.db.Read(r.Context(), id)
	if err != nil {
		if err == domain.ErrUserNotFound {
			http.Error(
				w,
				"user not found",
				http.StatusNotFound)
		} else {
			http.Error(
				w,
				"internal server error",
				http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(
			w,
			"failed to encode response",
			http.StatusInternalServerError)
	}
}

func (h *UserHandler) Save(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed)
		return
	}

	var user domain.User
	if err := json.NewDecoder(r.Body).
		Decode(&user); err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest)
		return
	}

	if err := h.db.Save(r.Context(), user); err != nil {
		http.Error(
			w,
			"failed to save user",
			http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
