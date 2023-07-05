package models

import (
	"time"
)

type User struct {
	ID        int       `form:"primary_key" json:"id"`
	FirstName string    `form:"not null" json:"first_name" binding:"required"`
	LastName  string    `form:"not null" json:"last_name" binding:"required"`
	PhoneNo   string    `form:"not null" json:"phone_no" binding:"len=10"`
	Email     string    `form:"not null;unique" json:"email" binding:"required,email"`
	Country   string    `form:"not null" json:"country" binding:"required"`
	Passcode  string    `form:"not null" json:"passcode" binding:"required,min=6"`
	CreatedAt time.Time `form:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `form:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
