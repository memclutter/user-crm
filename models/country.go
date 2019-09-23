package models

type Country struct {
	Code  string `gorm:"primary_key" json:"code"`
	Name  string `json:"name"`
	Users []User `gorm:"foreignkey:CountryCode" json:"-"`
}
