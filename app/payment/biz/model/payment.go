package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model
	UserId        uint32    `json:"user_id"`
	OrderId       string    `json:"order_id"`
	TransactionId string    `json:"transaction_id"`
	Amount        float32   `json:"amount"`
	PayAt         time.Time `json:"pay_at"`
}

func (Payment) TableName() string {
	return "payment"
}

func CreatePayment(ctx context.Context, db *gorm.DB, payment *Payment) error {
	return db.WithContext(ctx).Model(&Payment{}).Create(payment).Error

}
