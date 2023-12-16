package service

import (
	"ZAtest/models"
	"ZAtest/pkg/repository"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type NewsService struct {
	repo repository.News
}

func NewNewsService(repo repository.News) *NewsService {
	return &NewsService{repo: repo}
}
func (s *NewsService) GetAllNews() ([]models.News, error) {
	return s.repo.GetAllNews()
}

func (s *NewsService) EditNewsById(id int64, news models.News) error {
	// Валидация данных.
	if err := validator.New().Struct(news); err != nil {
		logrus.Debugf("Validation failed")
		return fmt.Errorf("Failed to validate data")
	}
	// Вызов метода репозитория для редактирования новости.
	return s.repo.EditNewsById(id, news)
}
