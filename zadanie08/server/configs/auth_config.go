package configs

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

var GoogleOauth2Config *oauth2.Config

func InitAuthConfig() {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	GoogleOauth2Config = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:1323/google/callback",
		Scopes: []string{
			"email",
			"profile",
		},
		Endpoint: google.Endpoint,
	}
}
