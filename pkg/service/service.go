package service

import (
	"ZAtest/models"
	"ZAtest/pkg/repository"
)

type Service struct {
	News
}
type News interface {
	GetAllNews() ([]models.News, error)
	EditNewsById(id int64, news models.News) error
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		News: NewNewsService(repos),
	}
}
