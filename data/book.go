package data

type Book struct {
	ArticleID uint   `json:"article_id" gorm:"primaryKey"`
	Title     string `json:"title" gorm:"type:varchar(100)"`
}

func (a Book) Create(data *Book) error {
	q := db.Create(data)
	err := q.Error
	if err != nil {
		return err
	}
	return nil
}
