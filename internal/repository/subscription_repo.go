package repository

import (
	"subminder/internal/domain"
	"time"

	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	Create(sub *domain.Subscription) error
	GetAll() ([]domain.Subscription, error)
	GetExpiringSubscriptions(days int) ([]domain.Subscription, error)
}

type subscriptionRepo struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &subscriptionRepo{db: db}
}

func (r *subscriptionRepo) Create(sub *domain.Subscription) error {
	return r.db.Create(sub).Error
}

func (r *subscriptionRepo) GetAll() ([]domain.Subscription, error) {
	var subscriptions []domain.Subscription
	err := r.db.Preload("Category").Find(&subscriptions).Error
	return subscriptions, err
}

func (r *subscriptionRepo) GetExpiringSubscriptions(days int) ([]domain.Subscription, error) {
	var subscriptions []domain.Subscription

	targetDate := time.Now().AddDate(0, 0, days)

	err := r.db.Where("active = ? AND renewal_date <= ?", true, targetDate).Find(&subscriptions).Error
	return subscriptions, err
}
