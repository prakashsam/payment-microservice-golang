package controllers

import (
	"paymentservice/services"

	"github.com/kataras/iris/v12"
)

func GetPayment(ctx iris.Context) {
	orderID := ctx.Params().Get("order_id")
	payment, err := services.GetPaymentByOrderID(orderID)
	if err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.WriteString("Payment not found for given order ID")
		return
	}

	ctx.JSON(payment)
}
