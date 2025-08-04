package services

import (
	"gorm.io/gorm"
	"paymentservice/models"
)

type PaymentService struct {
	DB *gorm.DB
}

func (ps *PaymentService) CreatePayment(orderID string, amount int) error {
	p := models.Payment{OrderID: orderID, Amount: amount, Status: "Processed"}
	return ps.DB.Create(&p).Error
}

func (ps *PaymentService) GetPaymentByOrderID(orderID string) (models.Payment, error) {
	var payment models.Payment
	err := ps.DB.Where("order_id = ?", orderID).First(&payment).Error
	return payment, err
}
