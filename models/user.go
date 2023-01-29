package models

import "gorm.io/gorm"

type User struct {
	Id             int       `json:"id" form:"id" gorm:"size:10;not null;uniqueIndex;primary_key;autoIncrement"`
	Fullname       string    `json:"fullname" form:"fullname" validate:"required" gorm:"size:30;not null"`
	Username       string    `json:"username" form:"username" validate:"required" gorm:"size:20;not null"`
	Email          string    `json:"email" form:"email" validate:"required,email" gorm:"size:40;not null;uniqueIndex"`
	Gender         string    `json:"gender" form:"gender" validate:"required" gorm:"size:10;not null"`
	Password       string    `json:"password" form:"password" validate:"required" gorm:"size:100;not null"`
	Status         string    `json:"status" form:"status" validate:"required" gorm:"size:10;not null;default:'user'"`
	Avatar         string    `json:"avatar" form:"avatar" gorm:"size:150;not null"`
	RefreshToken   string    `json:"refreshToken" form:"refreshToken" gorm:"size:5o"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {  
	if u.Avatar == "" {
		if u.Gender == "male" {
			u.Avatar = "https://www.pngall.com/wp-content/uploads/12/Avatar-Profile-PNG-Pic.png"
		} else {
			u.Avatar = "https://cdn1.iconfinder.com/data/icons/user-pictures/100/female1-512.png"
		}
	}

	return
}

type DisplayUserData struct {
	Avatar   string `form:"avatar"`
	Fullname string `form:"fullname" validate:"required"`
	Username string `form:"username" validate:"required"`
	Email    string `form:"email" validate:"required,email"`
	Gender   string `form:"gender" validate:"required"`
	Status   string `form:"status" validate:"required"`
}