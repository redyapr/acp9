package config

import (
	"acp9-redy-gigih/models/cart"
	"acp9-redy-gigih/models/category"
	"acp9-redy-gigih/models/product"
	"acp9-redy-gigih/models/transaction"
	"acp9-redy-gigih/models/user"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

func InitDB() {
	connectionString := Env("MYSQL_USER") + ":" + Env("MYSQL_PASS") + "@tcp(" + Env("MYSQL_HOST") + ":" + Env(
		"MYSQL_PORT") + ")/" + Env("MYSQL_NAME") + "?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	initialMigration()
}

func initialMigration() {
	var models = []interface{}{
		&user.User{},
		&category.Category{},
		&product.Product{},
		&cart.Cart{},
		&cart.CartDetail{},
		&transaction.Transaction{},
		&transaction.TransactionDetail{},
	}
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models...)
}

func InitDBTest() {
	connectionString := "root:@tcp(localhost:3306)/acp_test?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

func SetupEchoDB() *echo.Echo {
	InitDBTest()
	e := echo.New()
	return e
}
