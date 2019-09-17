package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username    string
	Email       string
	Birthday    time.Time
	CountryCode string
	Gender      string
}
