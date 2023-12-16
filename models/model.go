package models

//reform:news
type News struct {
	Id         int64   `json:"id" reform:"id"`
	Title      string  `json:"title" reform:"title" validate:"required"`
	Content    string  `json:"content" reform:"content" validate:"required"`
	Categories []int64 `json:"categories" reform:"-"`
}

type NewsCategories struct {
	NewsId     int64 `reform:"news_id"`
	CategoryId int64 `reform:"category_id"`
}
