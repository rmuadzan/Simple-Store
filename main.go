package main

import (
	"log"
	"simple-catalog-v2/connect"
	"simple-catalog-v2/controllers"
	"simple-catalog-v2/middlewares"
	"simple-catalog-v2/models"
	"simple-catalog-v2/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Renderer = repositories.NewRenderer("views/*.html", true)
	e.HTTPErrorHandler = repositories.ErrorHandler
	e.Validator = &models.CustomValidator{Validator: validator.New()}

	err := connect.MySqlConnect().AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	err = connect.MySqlConnect().AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal(err)
	}

	err = connect.MySqlConnect().AutoMigrate(&models.Images{})
	if err != nil {
		log.Fatal(err)
	}

	// err = PopulateProduct()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	e.Use(middlewares.MiddlewareContextValue)
	e.Use(middlewares.MiddlewareJWTAuthorization)
	e.Use(middlewares.MiddlewareLogging)

	e.GET("/", controllers.IndexHandler)
	e.GET("/login", controllers.LoginHandler)
	e.GET("/signup", controllers.SignUpPage)
	e.POST("/signup", controllers.SignUpHandler)
	e.POST("/auth", controllers.AuthUserHandler)
	e.GET("logout", controllers.LogoutHandler)
	e.GET("/products", controllers.AllProductsHandler)
	e.GET("/my-products", controllers.UserProductsHandler)
	e.POST("/my-products", controllers.CreateProductHandler)
	e.GET("/my-cart", nil)
	e.GET("/products/add", controllers.AddProductHandler)
	e.GET("/products/:id", controllers.GetProductByIdHandler)
	e.POST("/products/:id", controllers.UpdateProductByIdHandler)
	e.GET("/products/:id/edit", controllers.EditProductHandler)
	e.POST("/products/:id/delete", controllers.DeleteProductByIdHandler)
	e.GET("/search", controllers.SearchProductHandler)
	e.GET("/profile", controllers.GetUserInformationHandler)
	e.POST("/profile", controllers.UpdateUserInformationHandler)
	e.GET("/profile/edit", controllers.EditUserInformationHandler)
	e.GET("/about", controllers.AboutHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
