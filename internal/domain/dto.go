package domain

import "time"

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type CreateSubscriptionRequest struct {
	Name         string    `json:"name" binding:"required"`
	Price        float64   `json:"price" binding:"required,min=0"`
	Currency     string    `json:"currency" binding:"required,oneof=TRY USD EUR"`
	BillingCycle string    `json:"billing_cycle" binding:"oneof=Monthly Yearly"`
	StartDate    time.Time `json:"start_date"`
	RenewalDate  time.Time `json:"renewal_date"`
	CategoryID   uint      `json:"category_id" binding:"required"`
}
type SubscriptionResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	Currency     string    `json:"currency"`
	BillingCycle string    `json:"billing_cycle"`
	StartDate    time.Time `json:"start_date"`
	RenewalDate  time.Time `json:"renewal_date"`
	Active       bool      `json:"active"`
	CategoryID   uint      `json:"category_id"`
}

type CategoryResponse struct {
	ID            uint                   `json:"id"`
	Name          string                 `json:"name"`
	Description   string                 `json:"description"`
	Subscriptions []SubscriptionResponse `json:"subscriptions,omitempty"`
}

func ToSubscriptionResponse(sub Subscription) SubscriptionResponse {
	return SubscriptionResponse{
		ID:           sub.ID,
		Name:         sub.Name,
		Price:        sub.Price,
		Currency:     sub.Currency,
		BillingCycle: sub.BillingCycle,
		StartDate:    sub.StartDate,
		RenewalDate:  sub.RenewalDate,
		Active:       sub.Active,
		CategoryID:   sub.CategoryID,
	}
}

func ToSubscriptionResponseList(subs []Subscription) []SubscriptionResponse {
	var responses []SubscriptionResponse
	for _, s := range subs {
		responses = append(responses, ToSubscriptionResponse(s))
	}
	return responses
}

func ToCategoryResponse(cat Category) CategoryResponse {
	var subResponses []SubscriptionResponse

	if len(cat.Subscriptions) > 0 {
		for _, sub := range cat.Subscriptions {
			cleanSub := ToSubscriptionResponse(sub)
			subResponses = append(subResponses, cleanSub)
		}
	}

	return CategoryResponse{
		ID:            cat.ID,
		Name:          cat.Name,
		Description:   cat.Description,
		Subscriptions: subResponses,
	}
}

func ToCategoryResponseList(cats []Category) []CategoryResponse {
	var responses []CategoryResponse
	for _, s := range cats {
		responses = append(responses, ToCategoryResponse(s))
	}
	return responses
}
