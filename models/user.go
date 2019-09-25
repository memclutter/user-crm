package models

import (
	"time"
)

type User struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Birthday    time.Time `json:"birthday"`
	CountryCode string    `json:"countryCode"`
	Gender      string    `json:"gender"`
}
