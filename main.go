package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Role      string         `json:"role"`
	Status    string         `json:"status"`
	OTP       string         `json:"otp"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type Category struct {
	gorm.Model
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
}

type Product struct {
	gorm.Model
	ID         uint           `gorm:"primarykey" json:"id"`
	CategoryID int            `json:"categoryId"`
	Name       string         `json:"name"`
	Price      int            `json:"price"`
	Stockint   int            `json:"stock"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt"`
}

type Cart struct {
	gorm.Model
	ID        uint           `gorm:"primarykey" json:"id"`
	UserId    int            `json:"userId"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type CartDetail struct {
	gorm.Model
	ID        uint           `gorm:"primarykey" json:"id"`
	CartId    int            `json:"cartId"`
	ProductID int            `json:"productId"`
	Qty       int            `json:"qty"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type Transaction struct {
	gorm.Model
	ID                uint           `gorm:"primarykey" json:"id"`
	UserID            int            `json:"userId"`
	TransactionStatus string         `json:"status"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `json:"deletedAt"`
}

type TransactionDetail struct {
	gorm.Model
	ID            uint           `gorm:"primarykey" json:"id"`
	TransactionID int            `json:"transactionId"`
	ProductID     int            `json:"productId"`
	DetailQTY     int            `json:"qty"`
	DetailPrice   int            `json:"price"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `json:"deletedAt"`
}

func InitDB() {
	connectionString := "u1116242_acp:u1116242_acp@tcp(5.181.216.124:3306)/u1116242_acp?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Category{})
	DB.AutoMigrate(&Product{})
	DB.AutoMigrate(&Cart{})
	DB.AutoMigrate(&CartDetail{})
	DB.AutoMigrate(&Transaction{})
	DB.AutoMigrate(&TransactionDetail{})
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
