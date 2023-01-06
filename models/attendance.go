package models

import (
	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model
	User_id         int    `json:"user_id"`
	User            *User  `gorm:"foreignKey:UserId"`
	Type_attendance string `json:"type_attendance"`
}
