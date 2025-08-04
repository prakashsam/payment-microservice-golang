package routes

import (
	"paymentservice/controllers"
	"paymentservice/db"
	"paymentservice/services"

	"github.com/kataras/iris/v12"
)

func RegisterPaymentRoutes(app *iris.Application) {
	paymentService := &services.PaymentService{DB: db.GetDBConnection()}
	paymentController := &controllers.PaymentController{Service: paymentService}
	payment := app.Party("/payment")
	{
		payment.Get("/{order_id:string}", paymentController.GetPayment)
	}
}
