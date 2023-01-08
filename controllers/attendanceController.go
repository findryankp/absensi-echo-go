package controllers

import (
	"absensi/configs"
	"absensi/helpers"
	"absensi/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func isExistAttendance(attendance *models.Attendance, form models.Attendance) bool {
	query := configs.DB.Where(form).First(&attendance)

	if query.Error != nil {
		return false
	}

	return true
}

func CreateAttendance(c echo.Context) error {
	var attendance models.Attendance
	form := models.Attendance{
		UserId:           1,
		TypeAttendanceId: 1,
	}

	if isExist := isExistAttendance(&attendance, form); isExist {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "You have already check in", nil)
	}

	//create attendance
	configs.DB.Create(&form)

	return helpers.ResponseJson(c, http.StatusCreated, true, "You have already check in", nil)
}

func AttendaceRiwayat(c echo.Context) error {
	var attendance []models.Attendance
	configs.DB.Preload("TypeAttendance").Preload("User").Find(&attendance)
	return helpers.ResponseJson(c, http.StatusOK, true, "-", attendance)
}
