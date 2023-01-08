package controllers

import (
	"absensi/configs"
	"absensi/helpers"
	"absensi/models"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func IsExistAttendance(c echo.Context, typeAttendanceId int) bool {
	dt := time.Now()
	dateNow := dt.Format("2006-01-02") + "%"

	var attendance models.Attendance
	form := models.Attendance{
		UserId:           helpers.ClaimToken(c).ID,
		TypeAttendanceId: typeAttendanceId,
	}

	if isExist := configs.DB.Where(&form).Where("created_at LIKE ?", dateNow).First(&attendance); isExist.Error == nil {
		return false
	}

	return true
}

func CreateAttendance(c echo.Context) error {
	dt := time.Now()
	dateNow := dt.Format("2006-01-02") + "%"

	var attendance, form models.Attendance
	c.Bind(&form)
	form.UserId = helpers.ClaimToken(c).ID

	isExist := configs.DB.Where(&form).Where("created_at LIKE ?", dateNow).Preload("TypeAttendance").First(&attendance)
	if isExist.Error == nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, fmt.Sprintf("You have already %s at %s", attendance.TypeAttendance.Type, attendance.CreatedAt), nil)
	}

	//create attendance
	if create := configs.DB.Create(&form); create.Error != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "Error Input", create.Error.Error())
	}

	return helpers.ResponseJson(c, http.StatusCreated, true, "Successfully", nil)
}

func AttendaceRiwayat(c echo.Context) error {
	var attendance []models.Attendance
	configs.DB.Where("user_id = ?", helpers.ClaimToken(c).ID).
		Preload("TypeAttendance").
		Preload("User").
		Find(&attendance)
	return helpers.ResponseJson(c, http.StatusOK, true, "-", attendance)
}
