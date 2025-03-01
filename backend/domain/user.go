package domain

import (
	"time"
)

type User struct {
	ID        int      `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  *string  `json:"last_name,omitempty"`
	Username  *string  `json:"username,omitempty"`
	Meta      UserMeta `json:"meta"`
}

type UserMeta struct {
	UserID               int       `json:"user_id"`
	IsBot                *bool     `json:"is_bot,omitempty"`
	LanguageCode         *string   `json:"language_code,omitempty"`
	IsPremium            *bool     `json:"is_premium,omitempty"`
	IsActive             *bool     `json:"is_active,omitempty"`
	IsBackground         *bool     `json:"is_background,omitempty"`
	IsExpanded           *bool     `json:"is_expanded,omitempty"`
	ViewportHeight       *int      `json:"viewport_height,omitempty"`
	ViewportStableHeight *int      `json:"viewport_stable_height,omitempty"`
	Platform             *string   `json:"platform,omitempty"`
	LastActiveAt         time.Time `json:"last_active_at"`
	FirstSeenAt          time.Time `json:"first_seen_at"`
	SessionDuration      int       `json:"session_duration"`
	LaunchCount          int       `json:"launch_count"`
	DeviceResolution     *string   `json:"device_resolution,omitempty"`
	DevicePixelRatio     *float64  `json:"device_pixel_ratio,omitempty"`
	BrowserInfo          *string   `json:"browser_info,omitempty"`
}
