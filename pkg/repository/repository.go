package repository

import (
	"ZAtest/models"

	"gopkg.in/reform.v1"
)

type News interface {
	GetAllNews() ([]models.News, error)
	EditNewsById(id int64, news models.News) error
}

type Repository struct {
	News
}

func NewRepository(db *reform.DB) *Repository {
	return &Repository{
		News: NewNewsRepository(db),
	}
}
