package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Id        int `json:"id" bson:"id" form:"id" gorm:"size:10;not null;primary_key;autoIncrement"`
	User      User
	UserID    int `json:"userID" bson:"userID" form:"userID" gorm:"size:10;not null;index"`
	Product Product
	ProductID int `json:"productdID" bson:"productID" form:"productID" gorm:"size:10;not null;index"`
	Quantity int `json:"quantity" bson:"quantity" form:"quantity" gorm:"not null"`
	OrderDate time.Time `json:"orderDate" bson:"orderDate" form:"orderDate" gorm:"not null"`
	Status string `json:"status" bson:"status" form:"status" gorm:"size:10;not null"`
	TotalPrice float64 `json:"totalPrice" bson:"totalPrice" form:"totalPrice" gorm:"not null;type:decimal(16,2)"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {  
	var stock int
	if err = tx.Debug().Model(&Product{}).Select("stock").Where("id = ?", o.ProductID).Find(&stock).Error; err != nil {
		return err
	}
	if o.Quantity > stock || o.Quantity < 0 {
		err = errors.New("sorry, invalid quantity. this happens because the quantity inputed exceeds the existing stock or less than 0")
	}
	return
}
  
func (o *Order) AfterCreate(tx *gorm.DB) (err error) {
	if o.Status == "paid" {
		err = tx.Debug().Model(&Product{}).Where("id = ?", o.ProductID).Update("stock", gorm.Expr("stock - ?", o.Quantity)).Error
	}
	return
}