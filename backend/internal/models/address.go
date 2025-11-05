package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	User_Id uint   `json:"user_id"`
	HouseNo string `json:"house_no"`
	Street  string `json:"street"`
	Line1   string `json:"line1"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}
