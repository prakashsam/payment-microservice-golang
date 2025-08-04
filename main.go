package main

import (
	"context"
	"paymentservice/db"
	"paymentservice/pubsub"
	router "paymentservice/routes"
	"paymentservice/services"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	godotenv.Load()
	db.InitDBConnection()

	paymentService := &services.PaymentService{DB: db.GetDBConnection()}
	go pubsub.StartOrderSubscriber(context.Background(), paymentService)

	router.RegisterPaymentRoutes(app)

	app.Listen(":8083")
}
