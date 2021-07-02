package config

import (
	"acp9-redy-gigih/models"
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	connectionString := Env("MYSQL_USER_TEST") + ":" + Env("MYSQL_PASS_TEST") + "@tcp(" + Env("MYSQL_HOST") + ":" + Env("MYSQL_PORT") + ")/" + Env("MYSQL_NAME_TEST") + "?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

func InitMongo() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb+srv://" + Env("MONGO_USER") + ":" + Env("MONGO_PASS") + "@" + Env("MONGO_HOST") + "/" + Env("MONGO_NAME") + "?retryWrites=true&w=majority")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	return client.Database(Env("MONGO_NAME")), nil
}
