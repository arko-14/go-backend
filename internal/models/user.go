package models

import "time"

// Request payload for Create/Update
type UserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

// Response payload
type UserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age,omitempty"` // Calculated field
}

// Logic: Calculate Age dynamically
func CalculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	// Subtract 1 if the birthday hasn't happened yet this year
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}