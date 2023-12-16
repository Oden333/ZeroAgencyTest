package repository

import (
	"ZAtest/models"
	"fmt"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"
)

type NewsRepository struct {
	db *reform.DB
}

func NewNewsRepository(db *reform.DB) *NewsRepository {
	return &NewsRepository{db: db}
}

func (r *NewsRepository) GetAllNews() ([]models.News, error) {
	var newsList []models.News
	query := `SELECT * FROM News`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var news models.News
		if err := rows.Scan(&news.Id, &news.Title, &news.Content); err != nil {
			return nil, err
		}
		newsList = append(newsList, news)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return newsList, nil
}

func (r *NewsRepository) EditNewsById(id int64, news models.News) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Проверяем существование новости с указанным ID
	var count int
	if err := tx.QueryRow("SELECT COUNT(*) FROM news WHERE id = $1", id).Scan(&count); err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("Новости с ID %d не найдены", id)
	}

	// Выполняем обновление новости в транзакции
	query := "UPDATE news SET title = $1, content = $2 WHERE id = $3"
	_, err = tx.Exec(query, news.Title, news.Content, id)
	if err != nil {
		logrus.Debug(err)
		return err
	}

	/*

			// Выполняем выборку категорий, которые надо удалить из таблицы связей
		   	var deleteSlice []int64
		   	query = "SELECT categoryid FROM newscategories WHERE id = $1 and category not in ($2)"
		   	rows, err := tx.Query(query)
		   	if err != nil {
		   		return err
		   	}
		   	defer rows.Close()
		   	for rows.Next() {
		   		var i int64
		   		if err := rows.Scan(&i); err != nil {
		   			return err
		   		}
		   		deleteSlice = append(deleteSlice, i)
		   	}

		   	// Удаляем связи между статьей и категориями
		   	_, err = tx.Exec("DELETE FROM newscategories WHERE id=$1 and categoryid in ($2)", id, pq.Array(&deleteSlice))
		   	if err != nil {
		   		return err
		   	}
	*/

	// Удаляем связи между статьей и категориями
	_, err = tx.Exec("DELETE FROM newscategories WHERE newsid=$1", id)
	if err != nil {
		return err
	}

	// Добавляем категории, которых нет в таблице связей
	query = `
        INSERT INTO newscategories (newsid, categoryid)
        VALUES ($1, unnest($2::bigint[]))
        ON CONFLICT (newsid, categoryid)
        DO NOTHING
    `
	_, err = tx.Exec(query, id, pq.Array(news.Categories))
	if err != nil {
		return err
	}

	// Завершаем транзакцию
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
