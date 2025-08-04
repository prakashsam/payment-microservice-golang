package controllers

import (
	"paymentservice/services"

	"github.com/kataras/iris/v12"
)

type PaymentController struct {
	Service *services.PaymentService
}

func (c *PaymentController) GetPayment(ctx iris.Context) {
	orderID := ctx.Params().Get("order_id")
	payment, err := c.Service.GetPaymentByOrderID(orderID)
	if err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.WriteString("Payment not found for given order ID")
		return
	}
	ctx.JSON(payment)
}
