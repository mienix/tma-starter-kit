package middlewares

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func TelegramAuthMiddleware(
	botToken string,
	next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		initData := r.Header.Get("X-Telegram-Init-Data")
		if initData == "" ||
			!Verify(initData, botToken) {
			http.Error(
				w,
				"invalid or missing init data",
				http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func Verify(initData, botToken string) bool {
	var hash string
	dataCheck := []string{}

	for _, param := range strings.Split(initData, "&") {
		if strings.HasPrefix(param, "hash=") {
			hash = strings.TrimPrefix(param, "hash=")
		} else {
			dataCheck = append(dataCheck, param)
		}
	}

	for i, param := range dataCheck {
		parts := strings.SplitN(param, "=", 2)
		if len(parts) == 2 {
			dataCheck[i] = parts[0] + "=" + mustUnescape(parts[1])
		}
	}

	sort.Strings(dataCheck)
	checkString := strings.Join(dataCheck, "\n")

	secretKey := hmacSHA256([]byte("WebAppData"), []byte(botToken))

	computedHash := hmacSHA256(secretKey, []byte(checkString))

	expectedHash, err := hex.DecodeString(hash)
	if err != nil {
		return false
	}

	return bytes.Equal(computedHash, expectedHash)
}

func hmacSHA256(key, data []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func mustUnescape(s string) string {
	res, _ := url.QueryUnescape(s)
	return res
}
