package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID      string `json:"id" gorm:"primaryKey;column:payment_id"`
	OrderID string `json:"order_id" gorm:"column:order_id"`
	Amount  int    `json:"amount"`
	Status  string `json:"status"`
}

func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = uuid.NewString()
	}
	return
}
