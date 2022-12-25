package repositories

import (
	"context"
	"simple-catalog-v2/connect"
	"simple-catalog-v2/models"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client = connect.ConnectMongo()
var colllection = client.Database("Store-api").Collection("products")

func getProductCount() (int, error) {
	count, err := colllection.CountDocuments(context.TODO(), bson.D{})
	return int(count), err
}

func getLastId() (int, error) {
	var product models.Product

	count, err := getProductCount()
	if err != nil {
		return count, err
	}

	cursor, err := colllection.Find(context.TODO(), bson.D{}, options.Find().SetSkip(int64(count)-1))
	if err != nil {
		return 0, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		if err := cursor.Decode(&product); err != nil {
			return 0, err
		}
	}

	if err := cursor.Err();  err != nil {
		return 0, err
	}

	return product.Id, nil
}

func GetAllProducts() (*[]*models.Product, error) {
	var result []*models.Product

	cursor, err := colllection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var product models.Product

		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}

		result = append(result, &product)
	}

	if err := cursor.Err();  err != nil {
		return nil, err
	}

	return &result, nil
}

func GetUserProducts(userID int) (*[]*models.Product, error) {
	var result []*models.Product

	cursor, err := colllection.Find(context.TODO(), bson.D{{Key: "userID", Value: userID}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var product models.Product

		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}

		result = append(result, &product)
	}

	if err := cursor.Err();  err != nil {
		return nil, err
	}

	return &result, nil	
}

func GetProductById(id int) (models.Product, error) {
	var result models.Product

	if err := colllection.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}}).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

func CreateProduct(data *models.Product) (error) {
	data.Lock = new(sync.Mutex)
	data.Lock.Lock()
	defer data.Lock.Unlock()
	id, err := getLastId()
	if err != nil {
		return err
	}
	thumbnail := data.Thumbnail
	if thumbnail == "" {
		thumbnail = "https://propertywiselaunceston.com.au/wp-content/themes/property-wise/images/no-image.png"
	}
	document := bson.D{
		{Key: "id", Value:  (id+1)},
		{Key: "userID", Value: data.UserID},
		{Key: "title", Value: data.Title},
		{Key: "description", Value: data.Description},
		{Key: "price", Value: data.Price},
		{Key: "discountPercentage", Value: data.DiscountPercentage},
		{Key: "rating", Value: data.Rating},
		{Key: "stock", Value: data.Stock},
		{Key: "brand", Value: data.Brand},
		{Key: "category", Value: data.Category},
		{Key: "thumbnail", Value: thumbnail},
		{Key: "images", Value: data.Images},
		{Key: "fPrice", Value: data.FPrice},
	}
	_, err = colllection.InsertOne(context.TODO(), document)
	return err
}

func UpdateProductById(data *models.Product) (error) {
	thumbnail := data.Thumbnail
	if thumbnail == "" {
		thumbnail = "https://propertywiselaunceston.com.au/wp-content/themes/property-wise/images/no-image.png"
	}

	update := bson.D{{Key: "$set",
        Value: bson.D{
            {Key: "id", Value:  data.Id},
			{Key: "title", Value: data.Title},
			{Key: "description", Value: data.Description},
			{Key: "price", Value: data.Price},
			{Key: "discountPercentage", Value: data.DiscountPercentage},
			{Key: "rating", Value: data.Rating},
			{Key: "stock", Value: data.Stock},
			{Key: "brand", Value: data.Brand},
			{Key: "category", Value: data.Category},
			{Key: "thumbnail", Value: thumbnail},
			{Key: "images", Value: data.Images},
			{Key: "fPrice", Value: data.FPrice},
        },
    }}

	_, err := colllection.UpdateOne(context.TODO(), bson.D{{Key: "id", Value: data.Id}}, update)
	return err
}

func DeleteProductById(id int) error {
	_, err := colllection.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id}})
	return err
}

func SearchProductByTitle(title string) (*[]*models.Product, error) {
	if title == "" {
		products, err := GetAllProducts()
		return products, err
	}
	var result []*models.Product

	model := mongo.IndexModel{Keys: bson.D{{Key: "title",Value: "text"}}}
	_, err := colllection.Indexes().CreateOne(context.TODO(), model)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "$text",Value: bson.D{{Key: "$search",Value: title}}}}
	cursor, err := colllection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var product models.Product

		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}

		result = append(result, &product)
	}

	if err := cursor.Err();  err != nil {
		return nil, err
	}

	return &result, nil
}