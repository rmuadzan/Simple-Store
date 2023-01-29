package repositories

import (
	"fmt"
	"simple-catalog-v2/models"
	"time"

	"gorm.io/gorm/clause"
)

func GetUserOrders(userID int, page int, perPage int) (*[]*models.Order, int, error) {
	var result []*models.Order
	var count int64
	skip := (page - 1) * perPage

	err := db.Debug().Model(&models.Order{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Debug().Model(&models.Order{}).Where("user_id = ?", userID).Preload(clause.Associations).Omit("User.Password").Limit(perPage).Offset(skip).Find(&result).Error
	if err != nil {
		return nil, 0, err
	}

	return &result, int(count), nil
}

func CreateOrder(data *models.Order) error {
	data.TotalPrice = float64(data.Quantity) * data.Product.FPrice
	fmt.Println(data.Product.FPrice, data.Quantity)
	data.OrderDate = time.Now()
	err := db.Debug().Model(&models.Order{}).Omit(clause.Associations).Create(&data).Error
	return err
}

func GetOrderById(id int) (models.Order, error) {
	var result models.Order

	err := db.Debug().Model(&models.Order{}).Preload(clause.Associations).Preload("Product.Images").Where("id = ?", id).Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func DeleteOrderById(id int) error {
	err := db.Debug().Model(&models.Order{}).Delete(&models.Order{}, &id).Error
	return err
}

