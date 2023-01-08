package controllers

import (
	"absensi/configs"
	"absensi/helpers"
	"absensi/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func GetAllActivity(c echo.Context) error {
	var activities []models.Activity
	configs.DB.Find(&activities)
	return helpers.ResponseJson(c, http.StatusOK, true, "-", activities)
}

func GetActivity(c echo.Context) error {
	var activity models.Activity

	if isExist := configs.DB.Where("id = ?", c.Param("id")).First(&activity); isExist.Error != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", nil)
	}

	return helpers.ResponseJson(c, http.StatusOK, true, "-", activity)
}

func CreateActivity(c echo.Context) error {
	var activity models.Activity
	c.Bind(&activity)

	layoutFormat := "2006-01-02 15:04:05"
	value := activity.Date
	date, err := time.Parse(layoutFormat, value)
	if err != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "Error time format", err.Error)
	}

	activity.Date = date.Format(layoutFormat)

	if create := configs.DB.Create(&activity); create.Error != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "Error", create.Error)
	}

	return helpers.ResponseJson(c, http.StatusCreated, true, "-", activity)
}

func UpdateActivity(c echo.Context) error {
	var activity, updateForm models.Activity
	if isExist := configs.DB.Where("id = ?", c.Param("id")).First(&activity); isExist.Error != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", nil)
	}

	c.Bind(&updateForm)

	layoutFormat := "2006-01-02 15:04:05"
	value := updateForm.Date
	date, err := time.Parse(layoutFormat, value)
	if err != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "Error time format", err.Error)
	}

	updateForm.Date = date.Format(layoutFormat)

	configs.DB.Model(&activity).Updates(updateForm)

	return helpers.ResponseJson(c, http.StatusOK, true, "-", activity)
}

func DeleteActivity(c echo.Context) error {
	var activity models.Activity
	if isExist := configs.DB.Where("id = ?", c.Param("id")).First(&activity); isExist.Error != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", nil)
	}

	configs.DB.Delete(&activity)
	return helpers.ResponseJson(c, http.StatusOK, true, "Activity deleted", nil)
}

func RiwayatActivity(c echo.Context) error {
	var activity []models.Activity
	configs.DB.Where("date like ?", "%"+c.QueryParam("date")+"%").Find(&activity)
	return helpers.ResponseJson(c, http.StatusOK, true, "Activity deleted", nil)
}
