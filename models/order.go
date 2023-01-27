package models

import "time"

type Order struct {
	Id        int `json:"id" bson:"id" form:"id" gorm:"size:10;not null;primary_key;autoIncrement"`
	User      User
	UserID    int `json:"userID" bson:"userID" form:"userID" gorm:"size:10;not null;index"`
	Product Product
	ProductID int `json:"productdID" bson:"productID" form:"productID" gorm:"not null;index"`
	Quantity int `json:"quantity" bson:"quantity" form:"quantity" gorm:"not null;index"`
	OrderDate time.Time `json:"orderDate" bson:"orderDate" form:"orderDate" gorm:"not null"`
	Status string `json:"status" bson:"status" form:"status" gorm:"size:10;not null"`
	TotalPrice float64 `json:"totalPrice" bson:"totalPrice" form:"status" gorm:"not null"`
}