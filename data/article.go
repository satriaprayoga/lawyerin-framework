package data

import (
	"time"
)

type Article struct {
	ArticleID   int       `json:"article_id" gorm:"primary_key;auto_increment:true"`
	UserInput   string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit    string    `json:"user_edit" gorm:"type:varchar(20)"`
	Title       string    `json:"title" gorm:"type:text"`
	Description string    `json:"description" gorm:"type:text"`
	Creator     string    `json:"creator" gorm:"type:text"`
	Category    string    `json:"category" gorm:"type:varchar(100)"`
	Bidang      string    `json:"bidang" gorm:"type:varchar(100)"`
	SubBidang   string    `json:"sub_bidang" gorm:"type:varchar(100)"`
	Slug        string    `json:"slug" gorm:"type:text"`
	TimeInput   time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit    time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}

func (a *Article) Create(data *Article) error {
	query := db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}

	return nil
}
