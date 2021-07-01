package routes

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)

	eJwt := e.Group("/")
	eJwt.Use(middleware.JWT([]byte(config.Env("JWT_SECRET"))))
	eJwt.GET("categories", controllers.GetCategoriesController)
	eJwt.GET("products", controllers.GetProductsController)
	eJwt.GET("products/:categorySlug", controllers.GetProductsByCategoryController)
	eJwt.POST("cart", controllers.AddCartController)
	eJwt.GET("cart", controllers.GetCartController)
	eJwt.PUT("cart", controllers.UpdateCartController)
	eJwt.DELETE("cart/:cartId", controllers.DeleteCartController)
	eJwt.POST("checkout", controllers.CheckoutController)
	eJwt.POST("payment", controllers.PaymentController)

	return e
}
