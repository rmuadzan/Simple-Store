package models

type User struct {
	Id int 	`json:"id" form:"id" gorm:"size:10;not null;uniqueIndex;primary_key;autoIncrement"`
	Fullname string `json:"fullname" form:"fullname" validate:"required" gorm:"size:30;not null"`
	Username string `json:"username" form:"username" validate:"required" gorm:"size:20;not null"`
	Email string `json:"email" form:"email" validate:"required,email" gorm:"size:40;not null;uniqueIndex"`
	Gender string `json:"gender" form:"gender" validate:"required" gorm:"size:10;not null"`
	Password string `json:"password" form:"password" validate:"required" gorm:"size:100;not null"`
	Status string `json:"status" form:"status" validate:"required" gorm:"size:10;not null;default:'user'"`
	Avatar string `json:"avatar" form:"avatar" gorm:"size:150;not null"`
}

type DisplayUserData struct {
	Avatar string `form:"avatar"`
	Fullname string `form:"fullname" validate:"required"`
	Username string `form:"username" validate:"required"`
	Email string `form:"email" validate:"required,email"`
	Gender string `form:"gender" validate:"required"`
	Status string `form:"status" validate:"required"`
}