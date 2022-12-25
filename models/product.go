package models

import "sync"

type Product struct {
	Lock *sync.Mutex
	Id                 int      `json:"id" bson:"id" form:"id"`
	Title              string   `json:"title" bson:"title" form:"title"`
	Description        string   `json:"description" bson:"description" form:"description"`
	Price              float64  `json:"price" bson:"price" form:"price"`
	DiscountPercentage float64  `json:"discountPercentage" bson:"discountPercentage" form:"discountPercentage"`
	Rating             float64  `json:"rating" bson:"rating" form:"rating"`
	Stock              int      `json:"stock" bson:"stock" form:"stock"`
	Brand              string   `json:"brand" bson:"brand" form:"brand"`
	Category           string   `json:"category" bson:"category" form:"category"`
	Thumbnail          string   `json:"thumbnail" bson:"thumbnail" form:"thumbnail"`
	Images             []string `json:"images" bson:"images" form:"images"`
	FPrice             float64  `json:"fPrice" bson:"fPrice" form:"fPrice"`
	UserID			   int 		`json:"userID" bson:"userID" form:"userID"`
}

type DisplayProductData struct {
	Id int
	Thumbnail string
	Price float64
	FPrice float64
	DiscountPercentage float64
	Title string
	StoreName string
}