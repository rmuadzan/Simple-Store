package models

type User struct {
	Id int 	`json:"id" form:"id"`
	Avatar string `json:"avatar" form:"avatar"`
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Email string `json:"email" form:"email" validate:"required,email"`
	Gender string `json:"gender" form:"gender" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Status string `json:"status" form:"status" validate:"required"`
}

type DisplayUserData struct {
	Avatar string `form:"avatar"`
	Fullname string `form:"fullname" validate:"required"`
	Username string `form:"username" validate:"required"`
	Email string `form:"email" validate:"required,email"`
	Gender string `form:"gender" validate:"required"`
	Status string `form:"status" validate:"required"`
}