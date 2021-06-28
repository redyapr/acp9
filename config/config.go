package config

import (
	"acp9-redy-gigih/models/cart"
	"acp9-redy-gigih/models/category"
	"acp9-redy-gigih/models/product"
	"acp9-redy-gigih/models/transaction"
	"acp9-redy-gigih/models/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

func InitDB() {
	connectionString := env("MYSQL_USER") + ":" + env("MYSQL_PASS") + "@tcp(" + env("MYSQL_HOST") + ":" + env(
		"MYSQL_PORT") + ")/" + env("MYSQL_NAME") + "?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	initialMigration()
}

func initialMigration() {
	var models = []interface{}{&user.User{}, &category.Category{}, &product.Product{}, &cart.Cart{},
		&cart.CartDetail{}, &transaction.Transaction{}, &transaction.TransactionDetail{}}
	DB.AutoMigrate(models...)
}
