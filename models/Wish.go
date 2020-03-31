package models

import (
	"time"
)

type Wish struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Description string    `gorm:"size:255;not null;" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	WishRefer uint
}
