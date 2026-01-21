package service

import (
	"subminder/internal/domain"
	"subminder/internal/repository"
)

type SubscriptionService interface {
	CreateSubscription(sub *domain.Subscription) error
	GetAllSubscriptions() ([]domain.Subscription, error)
}

type subscriptionService struct {
	repository repository.SubscriptionRepository
}

func NewSubscriptionService(repo repository.SubscriptionRepository) SubscriptionService {
	return &subscriptionService{repository: repo}
}

func (s *subscriptionService) CreateSubscription(sub *domain.Subscription) error {
	if sub.RenewalDate.IsZero() {
		sub.RenewalDate = sub.StartDate.AddDate(0, 1, 0)
	}
	return s.repository.Create(sub)
}

func (s *subscriptionService) GetAllSubscriptions() ([]domain.Subscription, error) {
	return s.repository.GetAll()
}
