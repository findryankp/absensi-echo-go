package models

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	UserId   int    `json:"user_id"`
	User     *User  `gorm:"foreignKey:UserId"`
	Activity string `json:"activity"`
	Date     string `json:"date"`
}
