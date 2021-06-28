package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	ID       uint   `gorm:"primarykey" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"UNIQUE" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Status   string `json:"status"`
	OTP      string `json:"otp"`
}

type Category struct {
	gorm.Model
	ID          uint   `gorm:"primarykey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Product struct {
	gorm.Model
	ID         uint   `gorm:"primarykey" json:"id"`
	CategoryID int    `json:"categoryId"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stockint   int    `json:"stock"`
}

type Cart struct {
	gorm.Model
	ID     uint `gorm:"primarykey" json:"id"`
	UserId int  `json:"userId"`
}

type CartDetail struct {
	gorm.Model
	ID        uint `gorm:"primarykey" json:"id"`
	CartId    int  `json:"cartId"`
	ProductID int  `json:"productId"`
	Qty       int  `json:"qty"`
}

type Transaction struct {
	gorm.Model
	ID                uint   `gorm:"primarykey" json:"id"`
	UserID            int    `json:"userId"`
	TransactionStatus string `json:"status"`
}

type TransactionDetail struct {
	gorm.Model
	ID            uint `gorm:"primarykey" json:"id"`
	TransactionID int  `json:"transactionId"`
	ProductID     int  `json:"productId"`
	DetailQTY     int  `json:"qty"`
	DetailPrice   int  `json:"price"`
}

func env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

func InitDB() {
	connectionString := env("MYSQL_USER") + ":" + env("MYSQL_PASS") + "@tcp(" + env("MYSQL_HOST") + ":" + env("MYSQL_PORT") + ")/" + env("MYSQL_NAME") + "?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(
		&User{},
		&Category{},
		&Product{},
		&Cart{},
		&CartDetail{},
		&Transaction{},
		&TransactionDetail{},
	)
}

func init() {
	InitDB()
	InitialMigration()
}

func IndexController(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to ACP9 Redy Gigih")
}

func RegisterController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Register")
}

func ConfirmationController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Register Confirmation")
}

func LoginController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Login")
}

func GetCategoriesController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Categories")
}

func GetProductsController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Products")
}

func GetProductsByCategoryController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Products by Category")
}

func AddCartController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Add to Cart")
}

func GetCartController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Get Cart")
}

func UpdateCartController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Update Cart")
}

func DeleteCartController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Delete Cart")
}

func CheckoutController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Checkout")
}

func PaymentController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Payment")
}

func main() {
	e := echo.New()

	e.GET("/", IndexController)
	e.POST("/register", RegisterController)
	e.POST("/register/confirm/:userOTP", ConfirmationController)
	e.POST("/login", LoginController)
	e.GET("/categories", GetCategoriesController)
	e.GET("/products", GetProductsController)
	e.GET("/products/:categotySlug", GetProductsByCategoryController)
	e.POST("/cart", AddCartController)
	e.GET("/cart", GetCartController)
	e.PUT("/cart", UpdateCartController)
	e.DELETE("/cart/:cartId", DeleteCartController)
	e.POST("/checkout", CheckoutController)
	e.POST("/payment", PaymentController)

	e.Start(":6969")
}
