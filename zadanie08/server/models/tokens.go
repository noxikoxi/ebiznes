package models

import "time"

type Tokens struct {
	UserID             uint      `gorm:"primary_key"`
	GoogleAccessToken  string    `json:"googleAccessToken"`
	GoogleTokenExpires time.Time `json:"googleTokenExpires"`
	JWT                string    `json:"jwt"`
	GithubAccessToken  string    `json:"githubAccessToken"`
	GithubTokenExpires time.Time `json:"githubTokenExpires"`
}
