package user

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int    `json:"id" example:"1" format:"int64"`
	Type      string `json:"type" example:"user type"`
	Firstname string `json:"firstname" example:"First name"`
	Lastname  string `json:"lastname" example:"Last name"`
}

//  example
var (
	ErrNoRow       = errors.New("no rows in result set")
	ErrNameInvalid = errors.New("name is empty")
)

// AddUser example
type AddUser struct {
	Name string `json:"name" example:"user name"`
}

// Validation example
func (a AddUser) Validation() error {
	switch {
	case len(a.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

// UpdateUser example
type UpdateUser struct {
	Name string `json:"name" example:"user name"`
}

// Validation example
func (a UpdateUser) Validation() error {
	switch {
	case len(a.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

// UsersAll example
func UsersAll(q string) ([]User, error) {
	if q == "" {
		return users, nil
	}
	as := []User{}
	for k, v := range users {
		if q == v.Firstname {
			as = append(as, users[k])
		}
	}
	return as, nil
}

// UserOne example
func UserOne(id int) (User, error) {
	for _, v := range users {
		if id == v.ID {
			return v, nil
		}
	}
	return User{}, ErrNoRow
}

// Insert example
func (a User) Insert() (int, error) {
	userMaxID++
	a.ID = userMaxID
	a.Firstname = fmt.Sprintf("user_%d", userMaxID)
	users = append(users, a)
	return userMaxID, nil
}

// Delete example
func Delete(id int) error {
	for k, v := range users {
		if id == v.ID {
			users = append(users[:k], users[k+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user id=%d is not found", id)
}

// Update example
func (a User) Update() error {
	for k, v := range users {
		if a.ID == v.ID {
			users[k].Firstname = a.Firstname
			return nil
		}
	}
	return fmt.Errorf("user id=%d is not found", a.ID)
}

var userMaxID = 3
var users = []User{
	{ID: 1, Firstname: "user_1"},
}
