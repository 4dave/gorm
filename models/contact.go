package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	FirstName string
	LastName  string
	Address   string
	Email     string
	Phone     string
	City      string
	State     string
	Zip       string
	Map       string
}

type ContactBody struct {
	FirstName string
	LastName  string
	Address   string
	Email     string
	Phone     string
	City      string
	State     string
	Zip       string
}
