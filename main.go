package main

import (
	"context"
	"paymentservice/db"
	"paymentservice/pubsub"
	router "paymentservice/routes"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	godotenv.Load()
	db.InitDBConnection()

	go pubsub.StartOrderSubscriber(context.Background())

	router.RegisterPaymentRoutes(app)

	app.Listen(":8082")
}
