package worker

import (
	"log"
	"subminder/internal/repository"
	"time"
)

type RenewalWorker interface {
	Start()
	checkRenewals()
}

type renewalWorker struct {
	repo repository.SubscriptionRepository
}

func NewRenewalWorker(repo repository.SubscriptionRepository) RenewalWorker {
	return &renewalWorker{repo: repo}
}

func (w *renewalWorker) Start() {
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for {
			<-ticker.C
			log.Println("Checking renewals...")
			w.checkRenewals()
		}
	}()
}

func (w *renewalWorker) checkRenewals() {
	subscriptions, err := w.repo.GetExpiringSubscriptions(3)
	if err != nil {
		log.Println("Worker error: ", err)
		return
	}
	if len(subscriptions) == 0 {
		log.Println("There are no subscription payments near")
		return
	}
	for _, subscription := range subscriptions {
		// Mock Email Service
		log.Printf("Mail sended: %s payment is near! (%s)", subscription.Name, subscription.RenewalDate.Format("2006-01-02"))
	}
}
