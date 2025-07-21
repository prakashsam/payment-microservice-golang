package routes

import (
	"paymentservice/controllers"

	"github.com/kataras/iris/v12"
)

func RegisterPaymentRoutes(app *iris.Application) {
	payment := app.Party("/payment")
	{
		payment.Get("/{order_id:string}", controllers.GetPayment)
	}
}
