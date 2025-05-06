package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Email    string `gorm:"uniqueIndex"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"-"`
}
