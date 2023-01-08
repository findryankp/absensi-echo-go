package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	UserId   int       `json:"user_id"`
	User     *User     `gorm:"foreignKey:UserId"`
	Activity string    `json:"activity"`
	Date     time.Time `json:"date"`
}
