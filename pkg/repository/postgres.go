package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

const (
	newstable           = "news"
	newscategoriestable = "newscategories"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*reform.DB, error) {

	// Инициализация подключения к базе данных PostgreSQL
	sqlxDB, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = sqlxDB.Ping()
	if err != nil {
		return nil, err
	}
	// Инициализация connection pool с использованием reform
	db := reform.NewDB(sqlxDB.DB, postgresql.Dialect, nil)

	return db, nil
}
