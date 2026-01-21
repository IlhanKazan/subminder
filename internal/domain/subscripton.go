package domain

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	Name         string    `json:"name" gorm:"not null" binding:"required"`
	Price        float64   `json:"price" gorm:"not null" binding:"required,min=0"`
	Currency     string    `json:"currency" gorm:"default:'TRY'" binding:"required,oneof=TRY USD EUR"`
	BillingCycle string    `json:"billing_cycle" binding:"oneof=Monthly Yearly"`
	StartDate    time.Time `json:"start_date"`
	RenewalDate  time.Time `json:"renewal_date"`
	Active       bool      `json:"active" gorm:"default:true"`
	CategoryID   uint      `json:"category_id" binding:"required"`
	Category     *Category `json:"category" gorm:"foreignkey:CategoryID"`
}
