package users

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email    string `gorm:"index"`
	Password string
	Name     string
}
