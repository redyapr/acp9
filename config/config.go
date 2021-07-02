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
	connectionString := Env("MYSQL_USER") + ":" + Env("MYSQL_PASS") + "@tcp(" + Env("MYSQL_HOST") + ":" + Env("MYSQL_PORT") + ")/" + Env("MYSQL_NAME") + "?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	initialMigration()
}

func initialMigration() {
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Category{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Product{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Cart{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Transaction{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.TransactionDetail{})
}

func InitDBTest() {
	connectionString := Env("MYSQL_USER_TEST") + ":" + Env("MYSQL_PASS_TEST") + "@tcp(" + Env("MYSQL_HOST_TEST") + ":" + Env("MYSQL_PORT") + ")/" + Env("MYSQL_NAME_TEST") + "?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}
