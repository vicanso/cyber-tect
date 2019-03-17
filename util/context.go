package util

import (
	"github.com/vicanso/cod"

	"github.com/vicanso/cyber-tect/config"
)

// GetTrackID get track id
func GetTrackID(c *cod.Context) string {
	trackCookie := config.GetTrackKey()
	if trackCookie == "" {
		return ""
	}
	cookie, _ := c.Cookie(trackCookie)
	if cookie == nil {
		return ""
	}
	return cookie.Value
}
