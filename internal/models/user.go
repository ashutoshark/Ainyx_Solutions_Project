package models

import "time"

// request for create user
type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB  string `json:"dob" validate:"required"`
}

// request for update
type UpdateUserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB  string `json:"dob" validate:"required"`
}

// response without age
type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
}

// response with age included
type UserWithAge struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age"`
}

// calculates age from dob
func CalculateAge(dob time.Time) int {
	today := time.Now()

	age := today.Year() - dob.Year()

	// check if bday passed or not
	if today.Month() < dob.Month() {
		age = age - 1
	} else if today.Month() == dob.Month() && today.Day() < dob.Day() {
		age = age - 1
	}

	return age
}
