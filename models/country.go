package models

type Country struct {
	Code  string `gorm:"primary_key"`
	Title string
	Users []User `gorm:"foreignkey:CountryCode"`
}
