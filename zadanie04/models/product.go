package models

import "time"

type Product struct {
	ID         uint      `gorm:"primary_key;autoIncrement" json:"id"`
	Name       string    `gorm:"size:255;not null;unique" json:"name"`
	Price      float32   `gorm:"not null" json:"price"`
	CategoryID uint      `gorm:"not null" json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Category Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}
