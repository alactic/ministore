package userm

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type User struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

//  example
var (
	ErrNoRow           = errors.New("no rows in result set")
	ErrPasswordInvalid = errors.New("Password is required")
	ErrEmailInvalid    = errors.New("Valid Email is required")
	ErrNilInvalid      = errors.New("Field is required")
)

// Validation for new user fields
func Validation(a User) error {
	fmt.Println("user :: ", a)
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	switch {
	case len(strings.TrimSpace(a.Password)) == 0:
		return ErrPasswordInvalid
	case len(strings.TrimSpace(a.Email)) == 0 || !rxEmail.MatchString(a.Email):
		return ErrEmailInvalid
	default:
		return nil
	}
}

type UserDetails struct {
	ID        string `json:"id" binding:"required"`
	Type      string `json:"type"  binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname"  binding:"required"`
	Email     string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Token string `json:"token"  binding:"required"`
}
