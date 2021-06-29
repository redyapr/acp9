package routes

import (
	"acp9-redy-gigih/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", controllers.IndexController)
	e.POST("/register", controllers.RegisterController)
	e.POST("/register/confirm/:userOTP", controllers.ConfirmationController)
	e.POST("/login", controllers.LoginController)

	// eJwt := e
	// eJwt.Use(middleware.JWT([]byte(config.Env("JWT_SECRET"))))
	e.GET("/categories", controllers.GetCategoriesController)
	e.GET("/products", controllers.GetProductsController)
	e.GET("/products/:categotySlug", controllers.GetProductsByCategoryController)
	e.POST("/cart", controllers.AddCartController)
	e.GET("/cart", controllers.GetCartController)
	e.PUT("/cart", controllers.UpdateCartController)
	e.DELETE("/cart/:cartId", controllers.DeleteCartController)
	e.POST("/checkout", controllers.CheckoutController)
	e.POST("/payment", controllers.PaymentController)

	return e
}
