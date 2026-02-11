package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func init() {
	validate.RegisterValidation("password_complex", validatePasswordComplex)
}

func validatePasswordComplex(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return IsValidatePassword(password)
}

func ValidatePassword(password string) error {
	if !IsValidatePassword(password) {
		if len(password) < 8 {
			return errors.New("รหัสผ่านต้องมีอย่างน้อย 8 ตัวอักษร")
		}
		return errors.New("รหัสผ่านต้องมีตัวอักษรตัวใหญ่ ตัวอักษรตัวเล็ก ตัวเลข และอักขระพิเศษอย่างน้อยตัวละ 1")
	}
	return nil
}
