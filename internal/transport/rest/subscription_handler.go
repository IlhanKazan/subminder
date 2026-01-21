package rest

import (
	"net/http"
	"subminder/internal/domain"
	"subminder/internal/service"

	"github.com/gin-gonic/gin"
)

type SubscriptionHandler struct {
	service service.SubscriptionService
}

func NewSubscriptionHandler(service service.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{service: service}
}

// CreateSubscription godoc
// @Summary      Create a new subscription
// @Description  Adds a new subscription to the database with validation.
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        subscription  body      domain.CreateSubscriptionRequest true  "Subscription Data"
// @Success      201           {object}  domain.SubscriptionResponse
// @Failure      400           {object}  map[string]string
// @Failure      500           {object}  map[string]string
// @Router       /subscriptions [post]
func (h *SubscriptionHandler) CreateSubscription(c *gin.Context) {
	var req domain.CreateSubscriptionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"JSON format is wrong: ": err.Error()})
		return
	}

	sub := domain.Subscription{
		Name:         req.Name,
		Price:        req.Price,
		Currency:     req.Currency,
		BillingCycle: req.BillingCycle,
		StartDate:    req.StartDate,
		RenewalDate:  req.RenewalDate,
		CategoryID:   req.CategoryID,
	}

	if err := h.service.CreateSubscription(&sub); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Subscription created", "subscription": sub})
}

// GetAllSubscriptions godoc
// @Summary      List all subscriptions
// @Description  Retrieves a list of all active and passive subscriptions.
// @Tags         subscriptions
// @Produce      json
// @Success      200  {array}   domain.SubscriptionResponse
// @Failure      500  {object}  map[string]string
// @Router       /subscriptions [get]
func (h *SubscriptionHandler) GetAllSubscriptions(c *gin.Context) {
	subs, err := h.service.GetAllSubscriptions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get"})
		return
	}

	responseList := domain.ToSubscriptionResponseList(subs)
	c.JSON(http.StatusOK, gin.H{"subscriptions": responseList})
}
