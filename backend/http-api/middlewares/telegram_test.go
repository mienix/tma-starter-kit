package middlewares_test

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
	"testing"

	"github.com/devflex-pro/tma-starter-kit/backend/http-api/middlewares"
)

func generateValidHash(initData, botToken string) string {
	params := strings.Split(initData, "&")
	var dataCheck []string

	for _, param := range params {
		if !strings.HasPrefix(param, "hash=") {
			dataCheck = append(dataCheck, param)
		}
	}

	for i, param := range dataCheck {
		parts := strings.SplitN(param, "=", 2)
		if len(parts) == 2 {
			key, _ := url.QueryUnescape(parts[0])
			value, _ := url.QueryUnescape(parts[1])
			dataCheck[i] = key + "=" + value
		}
	}

	sort.Strings(dataCheck)
	checkString := strings.Join(dataCheck, "\n")

	secretKeyHmac := hmac.New(sha256.New, []byte("WebAppData"))
	secretKeyHmac.Write([]byte(botToken))
	secretKey := secretKeyHmac.Sum(nil)

	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(checkString))
	return hex.EncodeToString(h.Sum(nil))
}

func TestVerify(t *testing.T) {
	botToken := "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11"

	validInitData := "auth_date=1712345678&user=%7B%22id%22%3A12345%2C%22first_name%22%3A%22Alice%22%7D"
	validHash := generateValidHash(validInitData, botToken)
	validInitDataWithHash := validInitData + "&hash=" + validHash

	tests := []struct {
		name     string
		initData string
		expected bool
	}{
		{
			name:     "validh hash",
			initData: validInitDataWithHash,
			expected: true,
		},
		{
			name:     "bad hash",
			initData: validInitData + "&hash=0000000000000000000000000000000000000000000000000000000000000000",
			expected: false,
		},
		{
			name:     "missed hash=",
			initData: validInitData,
			expected: false,
		},
		{
			name:     "malformed data",
			initData: "auth_date=9999999999&user=%7B%22id%22%3A12345%2C%22first_name%22%3A%22Alice%22%7D&hash=" + validHash,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := middlewares.Verify(tt.initData, botToken)
			if result != tt.expected {
				t.Errorf("verify() for %s returned %v, expected %v", tt.name, result, tt.expected)
			}
		})
	}
}
