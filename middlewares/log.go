package middlewares

import (
	"acp9-redy-gigih/config"
	"context"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/mgo.v2"
)

type logs struct {
	Created time.Time `bson:"created"`
	Message string    `bson:"message"`
}

type MongoWriter struct {
	sess *mgo.Session
}

func (mw *MongoWriter) Write(p []byte) (n int, err error) {
	mongo, _ := config.InitMongo()
	// err = mongo.Collection("logs").InsertOne(bson.M{
	// 	"created":  time.Now(),
	// 	"messages": string(p),
	// })
	_, err = mongo.Collection("logs").InsertOne(context.Background(), logs{time.Now(), string(p)})
	if err != nil {
		return
	}
	return len(p), nil
}

func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	session, _ := mgo.Dial("mongodb+srv://" + config.Env("MONGO_USER") + ":" + config.Env("MONGO_PASS") + "@" + config.Env("MONGO_HOST") + "/" + config.Env("MONGO_NAME") + "?retryWrites=true&w=majority")
	mw := &MongoWriter{session}
	log.SetOutput(mw)
}
