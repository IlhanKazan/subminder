package service

import (
	"subminder/internal/domain"
	"subminder/internal/repository"
)

type CategoryService interface {
	CreateCategory(category domain.Category) error
	GetAllCategories() ([]domain.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) CreateCategory(category domain.Category) error {
	return s.repo.Create(&category)
}

func (s *categoryService) GetAllCategories() ([]domain.Category, error) {
	return s.repo.GetAll()
}
