package data

import (
	"fmt"
	"time"

	"github.com/satriaprayoga/lawyerin-framework/pkg/utils"
	"gorm.io/gorm"
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

type ArticleForm struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Creator     string `json:"creator"`
	Category    string `json:"category"`
	Bidang      string `json:"bidang"`
	SubBidang   string `json:"sub_bidang"`
}

type ArticleResult struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Creator     string  `json:"creator"`
	Category    string  `json:"category"`
	Bidang      string  `json:"bidang"`
	SubBidang   string  `json:"sub_bidang"`
	Rank        float64 `json:"rank"`
}

func (a *Article) BeforeCreate(tx *gorm.DB) (err error) {
	a.Slug = a.Category + " " + a.Bidang + " " + a.SubBidang
	return
}

func (a *Article) BeforeUpdate(tx *gorm.DB) (err error) {
	a.Slug = a.Category + " " + a.Bidang + " " + a.SubBidang
	return
}

func (a *Article) Create(data *Article) error {
	//data.Slug = data.Category + " " + data.Bidang + " " + data.SubBidang
	query := db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}

	return nil
}

func (a *Article) Update(ID int, data interface{}) error {
	var err error

	q := db.Model(&Article{}).Where("article_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (a *Article) Delete(ID int) error {
	var err error
	q := db.Where("article_id=?", ID).Delete(&Article{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Article) GetByID(ID int) (*Article, error) {
	var result = &Article{}
	query := db.Where("article_id=?", ID).Find(&result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *Article) TextSearch(term string) (result *[]ArticleResult, err error) {
	searchTerm := utils.FormatSearch(term)
	queryString := fmt.Sprintf(`select title, description, creator,category, bidang, sub_bidang, ts_rank(text_search, to_tsquery('indonesian','%s')) as rank
	from article
	where text_search @@ to_tsquery('indonesian','%s')
	order by rank desc`, searchTerm, searchTerm)
	query := db.Raw(queryString).Scan(&result)
	err = query.Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
