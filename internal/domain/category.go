package domain

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name          string         `json:"name" gorm:"unique;not null" binding:"required"`
	Description   string         `json:"description" gorm:"not null" binding:"required"`
	Subscriptions []Subscription `json:"subscription,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
