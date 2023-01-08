package models

import (
	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model
	UserId           int             `json:"user_id"`
	User             *User           `json:"user" gorm:"foreignKey:UserId"`
	TypeAttendanceId int             `json:"type_attendance_id"`
	TypeAttendance   *TypeAttendance `json:"type_attendance" gorm:"foreignKey:TypeAttendanceId"`
}
