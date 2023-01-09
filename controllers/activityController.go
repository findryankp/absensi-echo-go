package controllers

import (
	"absensi/configs"
	"absensi/helpers"
	"absensi/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// Get All Activity godoc
// @Summary Get All Activity.
// @Description Get a list of Activity.
// @Tags Activity
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : ' Bearer <insert_your_token_here>'"
// @Success 200
// @Router /activity [get]
func GetAllActivity(c echo.Context) error {
	var activities []models.Activity
	configs.DB.Where("user_id = ?", helpers.ClaimToken(c).ID).Find(&activities)
	return helpers.ResponseJson(c, http.StatusOK, true, "-", activities)
}

// Get Activity By Id godoc
// @Summary Get Activity By Id.
// @Description Get a Activity.
// @Tags Activity
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : ' Bearer <insert_your_token_here>'"
// @Success 200
// @Router /activity/:id [get]
func GetActivity(c echo.Context) error {
	var activity models.Activity

	if isExist := configs.DB.Where("id = ?", c.Param("id")).First(&activity); isExist.Error != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", nil)
	}

	if activity.UserId != helpers.ClaimToken(c).ID {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "You do not have authorization for this id", nil)
	}

	return helpers.ResponseJson(c, http.StatusOK, true, "-", activity)
}

// Create Activity godoc
// @Summary Create Activity .
// @Description Create Activity .
// @Tags Activity
// @Param Body body models.FormActivity true "the body to create a new Activity"
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : ' Bearer <insert_your_token_here>'"
// @Success 201
// @Router /activity [post]
func CreateActivity(c echo.Context) error {
	var activity models.Activity
	c.Bind(&activity)
	activity.UserId = helpers.ClaimToken(c).ID

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

// Update Activity godoc
// @Summary Update Activity .
// @Description Update Activity .
// @Tags Activity
// @Param Body body models.FormActivity true "the body to Update a new Activity"
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : ' Bearer <insert_your_token_here>'"
// @Success 200
// @Router /activity/:id [put]
func UpdateActivity(c echo.Context) error {
	var activity, updateForm models.Activity
	if isExist := configs.DB.Where("id = ?", c.Param("id")).First(&activity); isExist.Error != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", nil)
	}

	c.Bind(&updateForm)
	activity.UserId = helpers.ClaimToken(c).ID

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

// Delete Activity By Id godoc
// @Summary Delete Activity By Id.
// @Description Delete a Activity.
// @Tags Activity
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : ' Bearer <insert_your_token_here>'"
// @Success 200
// @Router /activity/:id [delete]
func DeleteActivity(c echo.Context) error {
	var activity models.Activity
	if isExist := configs.DB.Where("id = ?", c.Param("id")).First(&activity); isExist.Error != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", nil)
	}

	configs.DB.Delete(&activity)
	return helpers.ResponseJson(c, http.StatusOK, true, "Activity deleted", nil)
}

// Riwayat Activity godoc
// @Summary Riwayat Activity.
// @Description Riwayat Activity.
// @Tags Activity
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : ' Bearer <insert_your_token_here>'"
// @Success 200
// @Router /activity/riwayat [get]
func RiwayatActivity(c echo.Context) error {
	date := c.QueryParam("date") + "%"
	var activity []models.Activity
	if c.QueryParam("date") != "" {
		configs.DB.Where("user_id = ?", helpers.ClaimToken(c).ID).
			Where("date LIKE ?", date).
			Find(&activity)
	} else {
		configs.DB.Where("user_id = ?", helpers.ClaimToken(c).ID).Find(&activity)
	}

	return helpers.ResponseJson(c, http.StatusOK, true, "-", activity)
}
