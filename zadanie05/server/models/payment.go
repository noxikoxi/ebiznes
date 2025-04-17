package models

import "time"

type Payment struct {
	ID        uint      `gorm:"primary_key;autoIncrement" json:"id"`
	CartID    uint      `gorm:"not null" json:"cart_id"`
	Total     float32   `gorm:"not null" json:"total"`
	CreatedAt time.Time `json:"created_at"`

	Cart Cart `gorm:"foreignKey:CartID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"cart"`
}
