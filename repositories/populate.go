package repositories

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"simple-catalog-v2/models"
	"time"
)

type ResponseProduct struct {
	Product []models.Product `json:"products"`
	Images  [][]string       `json:"images"`
}

func PopulateProduct() error {
	res, err := http.Get("https://dummyjson.com/products?limit=100")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var data ResponseProduct

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	for _, product := range data.Product {
		product.UserID = rand.Intn(11-2) + 2
		// product.UserID = 1
		product.FPrice = product.Price * (100 - product.DiscountPercentage) / 100
		err := CreateProduct(&product)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func PopulateStore() error {
	var Stores []*models.User = []*models.User{
		{
			Fullname: "Toko Sumber Masalah",
			Username: "Tsm",
			Email: "tsm@email.com",
			Gender: "male",
			Password: "tokosumbermasalah",
			Status: "store",
		},
		{
			Fullname: "Toko Jaya Abadi",
			Username: "Tjy",
			Email: "tjy@email.com",
			Gender: "male",
			Password: "tokojayaabadi",
			Status: "store",
		},
		{
			Fullname: "Toko Sem Sul",
			Username: "Tss",
			Email: "tss@email.com",
			Gender: "female",
			Password: "tokosemsul",
			Status: "store",
		},
		{
			Fullname: "Urban Nitation Store",
			Username: "Uns",
			Email: "uns@email.com",
			Gender: "male",
			Password: "urbannitationstore",
			Status: "store",
		},
		{
			Fullname: "Planet Sud I&g",
			Username: "Psi",
			Email: "psi@email.com",
			Gender: "male",
			Password: "planetsudiang",
			Status: "store",
		},
		{
			Fullname: "Goob Loc Store",
			Username: "Gls",
			Email: "gls@email.com",
			Gender: "female",
			Password: "gooblocstore",
			Status: "store",
		},
		{
			Fullname: "Harapan Nil Store",
			Username: "Hns",
			Email: "hns@email.com",
			Gender: "female",
			Password: "harapannilstore",
			Status: "store",
		},
		{
			Fullname: "Gamma Store",
			Username: "Gs",
			Email: "gs@email.com",
			Gender: "male",
			Password: "gammastore",
			Status: "store",
		},
		{
			Fullname: "A Store",
			Username: "As",
			Email: "as@email.com",
			Gender: "male",
			Password: "astore",
			Status: "store",
		},
		{
			Fullname: "Setor Store",
			Username: "Ss",
			Email: "ss@email.com",
			Gender: "male",
			Password: "setorstore",
			Status: "store",
		},
	}

	for _, store := range Stores {
		err := CreateUser(store)
		return err
	}
	
	return nil
}