package models

import (
	"fmt"
	"strconv"
)

type Product struct {
	Id                 int `json:"-" bson:"id" form:"id" gorm:"size:10;not null;primary_key;autoIncrement"`
	User               User
	UserID             int      `json:"userID" bson:"userID" form:"userID" gorm:"size:10;index"`
	Title              string   `json:"title" bson:"title" form:"title" gorm:"size:200;not null"`
	Description        string   `json:"description" bson:"description" form:"description" gorm:"size:2000"`
	Price              float64  `json:"price" bson:"price" form:"price" gorm:"not null;type:decimal(16,2)"`
	DiscountPercentage float64  `json:"discountPercentage" bson:"discountPercentage" form:"discountPercentage" gorm:"default:0"`
	Rating             float64  `json:"rating" bson:"rating" form:"rating"`
	Stock              int      `json:"stock" bson:"stock" form:"stock" gorm:"not null;default:0"`
	Brand              string   `json:"brand" bson:"brand" form:"brand" gorm:"size:50;not null"`
	Category           string   `json:"category" bson:"category" form:"category" gorm:"size:20;not null"`
	Thumbnail          string   `json:"thumbnail" bson:"thumbnail" form:"thumbnail" gorm:"size:500;default:'https://propertywiselaunceston.com.au/wp-content/themes/property-wise/images/no-image.png'"`
	ImagesTemp         []string `json:"images" form:"images" gorm:"-"`
	Images             []Images `json:"-"`
	FPrice             float64  `json:"fPrice" bson:"fPrice" form:"fPrice" gorm:"not null;type:decimal(16,2)"`
}

func (p *Product) Init() {
	fPrice := p.Price * (100 - p.DiscountPercentage) / 100
	fPrice_string := fmt.Sprintf("%.2f", fPrice)
	p.FPrice, _ = strconv.ParseFloat(fPrice_string, 64)
	for _, img := range p.ImagesTemp {
		p.Images = append(p.Images, Images{Link: img})
	}
}

type DisplayProductData struct {
	Id                 int
	Thumbnail          string
	Price              float64
	FPrice             float64
	DiscountPercentage float64
	Title              string
	StoreName          string
}

type Images struct {
	Id        int `json:"id" form:"imagesID" gorm:"size:10;not null;primary_key;autoIncrement"`
	Product   Product
	ProductID int    `gorm:"size:10;index"`
	Link      string `json:"images" gorm:"size:500"`
}