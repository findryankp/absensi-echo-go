package models

type FormRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FormLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FormCreateAttendance struct {
	TypeAttendanceId int `json:"type_attendance_id"`
}

type FormActivity struct {
	Activity string `json:"activity"`
	Date     string `json:"date"`
}
