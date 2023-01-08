package models

import "gorm.io/gorm"

type TypeAttendance struct {
	gorm.Model
	Type string `json:"type"`
}
