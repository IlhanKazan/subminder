package rest

import (
	"net/http"
	"subminder/internal/domain"
	"subminder/internal/service"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// CreateCategory godoc
// @Summary      Create a new category for subscriptions
// @Description  Adds a new category to the database with validation.
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category  body      domain.CreateCategoryRequest  true  "Category Data"
// @Success      201           {object}  domain.CategoryResponse
// @Failure      400           {object}  map[string]string
// @Failure      500           {object}  map[string]string
// @Router       /categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req domain.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad Request: ": err.Error()})
		return
	}
	cat := domain.Category{
		Name:        req.Name,
		Description: req.Description,
	}
	if err := h.service.CreateCategory(cat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "There's an error while creating a category..."})
		return
	}
	response := domain.ToCategoryResponse(cat)
	c.JSON(http.StatusCreated, gin.H{"category": response})
}

// GetAllCategories godoc
// @Summary      List all categories
// @Description  Retrieves a list of all categories for subscriptions.
// @Tags         categories
// @Produce      json
// @Success      200  {array}   domain.CategoryResponse
// @Failure      500  {object}  map[string]string
// @Router       /categories [get]
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "There's an error while getting all categories..."})
		return
	}
	responseList := domain.ToCategoryResponseList(categories)
	c.JSON(http.StatusOK, gin.H{"categories": responseList})
}
