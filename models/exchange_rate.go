package models

import "time"

type ExchangeRate struct {
	Id           uint      `gorm:"primarykey" json:"_id"`
	FromCurrency string    `gorm:"from_currency" json:"fromCurrency" binding:"required"`
	ToCurrency   string    `gorm:"to_currency" json:"toCurrency" binding:"required"`
	Rate         float64   `gorm:"rate" json:"rate" binding:"required"`
	Date         time.Time `gorm:"date" json:"date"`
}
