package models

import "time"

type Cart struct {
	ID        uint      `gorm:"primary_key;autoIncrement" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	CartItems []CartItem `gorm:"foreignKey:CartID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"cart_items"`
}
