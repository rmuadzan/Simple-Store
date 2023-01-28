package repositories

import (
	"simple-catalog-v2/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	// "sync"
)

// func getProductCount() (int, error) {
// 	count, err := colllection.CountDocuments(context.TODO(), bson.D{})
// 	return int(count), err
// }

// func getLastId() (int, error) {
// 	var product models.Product

// 	count, err := getProductCount()
// 	if err != nil {
// 		return count, err
// 	}

// 	cursor, err := colllection.Find(context.TODO(), bson.D{}, options.Find().SetSkip(int64(count)-1))
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer cursor.Close(context.TODO())

// 	for cursor.Next(context.TODO()) {
// 		if err := cursor.Decode(&product); err != nil {
// 			return 0, err
// 		}
// 	}

// 	if err := cursor.Err();  err != nil {
// 		return 0, err
// 	}

// 	return product.Id, nil
// }

func GetAllProducts(page int, perPage int) (*[]*models.Product, int, error) {
	var result []*models.Product
	var count int64

	skip := (page - 1) * perPage
	// count, err := colllection.CountDocuments(context.TODO(), bson.D{})
	// if err != nil {
	// 	return nil, 0, err
	// }

	// cursor, err := colllection.Find(context.TODO(), bson.D{{}}, options.Find().SetSkip(int64(skip)).SetLimit(int64(perPage)))
	// if err != nil {
	// 	return nil, 0, err
	// }
	// defer cursor.Close(context.TODO())

	// for cursor.Next(context.TODO()) {
	// 	var product models.Product

	// 	if err := cursor.Decode(&product); err != nil {
	// 		return nil, 0, err
	// 	}

	// 	result = append(result, &product)
	// }

	// if err := cursor.Err();  err != nil {
	// 	return nil, 0, err
	// }

	err := db.Debug().Model(&models.Product{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Debug().Model(&models.Product{}).Limit(perPage).Offset(skip).Find(&result).Error

	return &result, int(count), err
}

func GetUserProducts(userID int) (*[]*models.Product, error) {
	var result []*models.Product

	// cursor, err := colllection.Find(context.TODO(), bson.D{{Key: "userID", Value: userID}})
	// if err != nil {
	// 	return nil, err
	// }
	// defer cursor.Close(context.TODO())

	// for cursor.Next(context.TODO()) {
	// 	var product models.Product

	// 	if err := cursor.Decode(&product); err != nil {
	// 		return nil, err
	// 	}

	// 	result = append(result, &product)
	// }

	// if err := cursor.Err();  err != nil {
	// 	return nil, err
	// }

	err := db.Debug().Model(&models.Product{}).Where("user_id = ?", userID).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil	
}

func GetProductById(id int) (models.Product, error) {
	var result models.Product

	// if err := colllection.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}}).Decode(&result); err != nil {
	// 	return result, err
	// }
	err := db.Debug().Model(&models.Product{}).Preload(clause.Associations).Where("id = ?", id).Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func CreateProduct(data *models.Product) (error) {
	// data.Lock = new(sync.Mutex)
	// data.Lock.Lock()
	// defer data.Lock.Unlock()
	// id, err := getLastId()
	// if err != nil {
	// 	return err
	// }
	// thumbnail := data.Thumbnail
	// if thumbnail == "" {
	// 	thumbnail = "https://propertywiselaunceston.com.au/wp-content/themes/property-wise/images/no-image.png"
	// }
	// document := bson.D{
	// 	{Key: "id", Value:  (id+1)},
	// 	{Key: "userID", Value: data.UserID},
	// 	{Key: "title", Value: data.Title},
	// 	{Key: "description", Value: data.Description},
	// 	{Key: "price", Value: data.Price},
	// 	{Key: "discountPercentage", Value: data.DiscountPercentage},
	// 	{Key: "rating", Value: data.Rating},
	// 	{Key: "stock", Value: data.Stock},
	// 	{Key: "brand", Value: data.Brand},
	// 	{Key: "category", Value: data.Category},
	// 	{Key: "thumbnail", Value: thumbnail},
	// 	{Key: "images", Value: data.Images},
	// 	{Key: "fPrice", Value: data.FPrice},
	// }
	// _, err = colllection.InsertOne(context.TODO(), document)
	err := db.Debug().Model(&models.Product{}).Create(&data).Error
	return err
}

func UpdateProductById(data *models.Product) (error) {
	// thumbnail := data.Thumbnail
	// if thumbnail == "" {
	// 	thumbnail = "https://propertywiselaunceston.com.au/wp-content/themes/property-wise/images/no-image.png"
	// }

	// update := bson.D{{Key: "$set",
    //     Value: bson.D{
    //         {Key: "id", Value:  data.Id},
	// 		{Key: "title", Value: data.Title},
	// 		{Key: "description", Value: data.Description},
	// 		{Key: "price", Value: data.Price},
	// 		{Key: "discountPercentage", Value: data.DiscountPercentage},
	// 		{Key: "rating", Value: data.Rating},
	// 		{Key: "stock", Value: data.Stock},
	// 		{Key: "brand", Value: data.Brand},
	// 		{Key: "category", Value: data.Category},
	// 		{Key: "thumbnail", Value: thumbnail},
	// 		{Key: "images", Value: data.Images},
	// 		{Key: "fPrice", Value: data.FPrice},
    //     },
    // }}

	// _, err := colllection.UpdateOne(context.TODO(), bson.D{{Key: "id", Value: data.Id}}, update)
	err := db.Debug().Model(&models.Product{}).Where("id = ?", data.Id).Updates(&data).Error
	return err
}

func DeleteProductById(id int) error {
	// _, err := colllection.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id}})
	err := db.Debug().Model(&models.Product{}).Select(clause.Associations).Delete(&models.Product{}, &id).Error
	return err
}

func SearchProductByTitle(title string) (*[]*models.Product, error) {
	// if title == "" {
	// 	products, err := GetAllProducts()
	// 	return products, err
	// }
	var result []*models.Product

	// model := mongo.IndexModel{Keys: bson.D{{Key: "title",Value: "text"}}}
	// _, err := colllection.Indexes().CreateOne(context.TODO(), model)
	// if err != nil {
	// 	return nil, err
	// }

	// filter := bson.D{{Key: "$text",Value: bson.D{{Key: "$search",Value: title}}}}
	// cursor, err := colllection.Find(context.TODO(), filter)
	// if err != nil {
	// 	return nil, err
	// }
	// defer cursor.Close(context.TODO())

	// for cursor.Next(context.TODO()) {
	// 	var product models.Product

	// 	if err := cursor.Decode(&product); err != nil {
	// 		return nil, err
	// 	}

	// 	result = append(result, &product)
	// }

	// if err := cursor.Err();  err != nil {
	// 	return nil, err
	// }
	
	titleLike := "%" + title + "%"
	err := db.Debug().Model(&models.Product{}).Where("title LIKE ?", titleLike).Find(&result).Error

	return &result, err
}

func UpdateProductStock(id int, quantity int) error {
	err := db.Debug().Model(&models.Product{}).Where("id = ?", id).Update("stock", gorm.Expr("stock - ?", quantity)).Error
	return err
}

func GetProductStock(id int) (int, error) {
	var result int
	err := db.Debug().Model(&models.Product{}).Select("stock").Where("id = ?", id).Find(&result).Error
	return result, err
}