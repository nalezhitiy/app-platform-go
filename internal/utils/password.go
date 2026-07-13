package utils

import (
	"errors"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPasswordTooShort = errors.New("password must be at least 8 characters long")
	ErrPasswordTooLong  = errors.New("password must not exceed 128 characters")
	ErrPasswordNoLower  = errors.New("password must contain at least one lowercase letter (a-z)")
	ErrPasswordNoUpper  = errors.New("password must contain at least one uppercase letter (A-Z)")
	ErrPasswordNoDigit  = errors.New("password must contain at least one digit (0-9)")
	ErrPasswordNoSpec   = errors.New("password must contain at least one special character")
)

// ValidatePassword перевіряє пароль на відповідність корпоративним стандартам безпеки.
// Довжина: 8-128 символів. Обов'язково: a-z, A-Z, 0-9, спецсимвол.
func ValidatePassword(password string) error {
	// 1. Перевірка довжини за кількістю рун (аби коректно рахувати UTF-8 символи, якщо вони будуть)
	length := 0
	var hasLower, hasUpper, hasDigit, hasSpecial bool

	for _, r := range password {
		length++
		switch {
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			// IsPunct покриває такі символи як ! . , ; : ? etc.
			// IsSymbol покриває такі символи як $ + < = > ^ ` ~ | etc.
			hasSpecial = true
		}
	}

	if length < 8 {
		return ErrPasswordTooShort
	}
	if length > 128 {
		return ErrPasswordTooLong
	}
	if !hasLower {
		return ErrPasswordNoLower
	}
	if !hasUpper {
		return ErrPasswordNoUpper
	}
	if !hasDigit {
		return ErrPasswordNoDigit
	}
	if !hasSpecial {
		return ErrPasswordNoSpec
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
