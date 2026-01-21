package database

import (
	"gorm.io/gorm"
	"log"
	"subminder/internal/domain"
	"time"
)

func SeedDatabase(db *gorm.DB) {
	var count int64
	db.Model(&domain.Category{}).Count(&count)

	if count > 0 {
		log.Println("Database already has mock data, skipping seeding process.")
		return
	}

	log.Println("Database is empty, inserting mock data...")

	entertainment := domain.Category{Name: "Entertainment", Description: "Movie, Series, Game"}
	work := domain.Category{Name: "Work & Software", Description: "Server, IDE, Cloud"}
	music := domain.Category{Name: "Music", Description: "Spotify, Apple Music"}

	subs := []domain.Subscription{
		{
			Name: "Netflix", Price: 229.99, Currency: "TRY", BillingCycle: "Monthly",
			StartDate:   time.Now().AddDate(0, -1, 0),
			RenewalDate: time.Now().AddDate(0, 0, 1),
			Category:    &entertainment,
		},
		{
			Name: "JetBrains", Price: 14.90, Currency: "USD", BillingCycle: "Yearly",
			StartDate:   time.Now().AddDate(0, -6, 0),
			RenewalDate: time.Now().AddDate(0, 0, 2),
			Category:    &work,
		},
		{
			Name: "Spotify", Price: 59.99, Currency: "TRY", BillingCycle: "Monthly",
			StartDate:   time.Now(),
			RenewalDate: time.Now().AddDate(0, 1, 0),
			Category:    &music,
		},
	}

	if err := db.Create(&subs).Error; err != nil {
		log.Println("Seed Error:", "error", err)
	} else {
		log.Println("Mock data successfully inserted!")
	}
}
