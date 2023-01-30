package main

import (
	"flag"
	"log"
	"os"
	"simple-catalog-v2/connect"
	"simple-catalog-v2/controllers"
	"simple-catalog-v2/middlewares"
	"simple-catalog-v2/models"
	"simple-catalog-v2/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/urfave/cli"
)

func runServer() {
	e := echo.New()
	e.Renderer = repositories.NewRenderer("views/*.html", true)
	e.HTTPErrorHandler = repositories.ErrorHandler
	e.Validator = &models.CustomValidator{Validator: validator.New()}

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
	e.GET("/products/add", controllers.AddProductHandler)
	e.GET("/products/:id", controllers.GetProductByIdHandler)
	e.POST("/products/:id", controllers.UpdateProductByIdHandler)
	e.GET("/products/:id/edit", controllers.EditProductHandler)
	e.POST("/products/:id/delete", controllers.DeleteProductByIdHandler)
	e.GET("/products/:id/order", controllers.AddOrderProduct)
	e.GET("/my-products", controllers.UserProductsHandler)
	e.POST("/my-products", controllers.CreateProductHandler)
	e.GET("/my-orders", controllers.UserOrderHandler)
	e.POST("/my-orders", controllers.OrderProductHandler)
	e.GET("/my-orders/:id", controllers.GetOrderByIdHandler)
	e.POST("/my-orders/:id/delete", controllers.DeleteOrderByIdHandler)
	e.GET("/search", controllers.SearchProductHandler)
	e.GET("/profile", controllers.GetUserInformationHandler)
	e.POST("/profile", controllers.UpdateUserInformationHandler)
	e.GET("/profile/edit", controllers.EditUserInformationHandler)
	e.GET("/about", controllers.AboutHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func main() {
	// CLI COMMAND
	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		cmdApp := cli.NewApp()
		cmdApp.Commands = []cli.Command{
			{
				Name: "db:migrate",
				Action: func (c *cli.Context) error {
					err := connect.MySqlConnect().AutoMigrate(&models.User{}, &models.Product{}, &models.Images{}, &models.Order{})
					if err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
			{
				Name: "db:populate",
				Action: func (c *cli.Context) error {
					err := repositories.PopulateStore()
					if err != nil {
						log.Fatal(err)
					}
					err = repositories.PopulateProduct()
					if err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
		}

		cmdApp.Run(os.Args)
	} else {
		runServer()
	}
}
