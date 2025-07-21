package services

import (
	"paymentservice/db"
	"paymentservice/models"
)

func CreatePayment(orderID string, amount int) {
	DB := db.GetDBConnection()
	p := models.Payment{OrderID: orderID, Amount: amount, Status: "Processed"}
	DB.Create(&p)
}

func GetPaymentByOrderID(orderID string) (models.Payment, error) {
	DB := db.GetDBConnection()
	var payment models.Payment
	err := DB.Where("order_id = ?", orderID).First(&payment).Error
	return payment, err
}
