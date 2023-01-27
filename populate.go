package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"simple-catalog-v2/models"
	"simple-catalog-v2/repositories"
	"time"
)

type Response struct {
	Product []models.Product `json:"products"`
	Images [][]string `json:"images"`
}

func PopulateProduct() error{
	res, err := http.Get("https://dummyjson.com/products?limit=100")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var data Response

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		panic(err)
	}


	rand.Seed(time.Now().UnixNano())
	for _, product := range data.Product {
		product.UserID = rand.Intn(11 - 2) + 2
		// product.UserID = 1
		product.FPrice = product.Price * (100 - product.DiscountPercentage) / 100
		product.Init()
		err := repositories.CreateProduct(&product)
		if err != nil {
			panic(err)
		}
	}

	return nil
}