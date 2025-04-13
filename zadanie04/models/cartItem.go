package models

import "time"

// CartItem
type CartItem struct {
	ID        uint      `gorm:"primary_key;autoIncrement" json:"id"`
	CartID    uint      `gorm:"not null" json:"cart_id"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	Quantity  uint      `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Product Product `gorm:"foreignKey:ProductID" json:"product"`
	Cart    Cart    `gorm:"foreignKey:CartID;constraint:OnUpdate:CASCADE,OnDelete:DELETE;" json:"cart"`
}
