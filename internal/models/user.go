package models

import "time"

type User struct {
	ID   int32     `json:"id"`
	Name string    `json:"name"`
	Dob  time.Time `json:"dob"`
	Age  int       `json:"age,omitempty"`
}

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required"`
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
}
