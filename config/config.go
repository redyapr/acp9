package config

import (
	"acp9-redy-gigih/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
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
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate([]interface{}{
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Cart{},
		&models.CartDetail{},
		&models.Transaction{},
		&models.TransactionDetail{},
	})
}

func InitDBTest() {
	connectionString := "root:@tcp(localhost:3306)/acp_test?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}
