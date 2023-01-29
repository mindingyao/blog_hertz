package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"user_name" column:"user_name"`
	Email    string `json:"email" column:"email"`
	Password string `json:"password" column:"password"`
}

func (u *User) TableName() string {
	return "users"
}