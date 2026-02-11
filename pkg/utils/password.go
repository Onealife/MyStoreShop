package utils

import (
	"errors"
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func validatePasswordComplecity(password string) error {

	if len(password) < 8 {
		return errors.New("รหัสผ่านต้องมีอย่างน้อย 8 ตัว")
	}

	var (
		hasUpper   = false
		hasLower   = false
		hashNumber = false
		hasSpecial = false
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hashNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("รหัสผ่านต้องมีตัวอักษรตัวใหญ่อย่างน้อย 1 ตัว")
	}
	if !hasLower {
		return errors.New("รหัสผ่านต้องมีตัวอักษรตัวเล็กอย่างน้อย 1 ตัว")
	}
	if !hashNumber {
		return errors.New("รหัสผ่านต้องมีตัวเลขอย่างน้อย 1 ตัว")
	}
	if !hasSpecial {
		return errors.New("รหัสผ่านต้องมีตัวอักขระพิเศษอย่างน้อย 1 ตัว")
	}

	return nil
}

func IsValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasLower, _ := regexp.MatchString(`[a-z]`, password)
	hasUpper, _ := regexp.MatchString(`[A-Z]`, password)
	hasNumber, _ := regexp.MatchString(`\d`, password)
	hasSpecial, _ := regexp.MatchString(`[!@#\$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~`+"`"+`]`, password)

	return hasLower && hasUpper && hasNumber && hasSpecial

}
