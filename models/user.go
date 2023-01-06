package models

import (
	"absensi/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;size:256;uniqueIndex"  valid:"required~Please insert username"`
	Email    string `json:"email" gorm:"uniqueIndex;not null;size:256" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `json:"password" gorm:"not null" form:"password" valid:"required~Your password is required,stringlength(6|100)~Minimal character of password is 6"`
	Age      int    `json:"age" gorm:"not null" valid:"range(8|150)"`
	Role     string `json:"role"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(user)

	if errCreate != nil {
		err = errCreate
		return
	}
	user.Password = helpers.EncryptPassword(user.Password)

	err = nil
	return
}
