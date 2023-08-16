package data

import (
	"fmt"
	"time"

	"github.com/satriaprayoga/lawyerin-framework/pkg/utils"
	"gorm.io/gorm"
)

type Peraturan struct {
	PeraturanID int       `json:"peraturan_id" gorm:"primary_key;auto_increment:true"`
	UserInput   string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit    string    `json:"user_edit" gorm:"type:varchar(20)"`
	Title       string    `json:"title" gorm:"type:text"`
	Description string    `json:"description" gorm:"type:text"`
	Creator     string    `json:"creator" gorm:"type:text"`
	Category    string    `json:"category" gorm:"type:varchar(100)"`
	Bidang      string    `json:"bidang" gorm:"type:varchar(100)"`
	SubBidang   string    `json:"sub_bidang" gorm:"type:varchar(100)"`
	Sejarah     string    `json:"sejarah" gorm:"type:text"`
	Dasar       string    `json:"dasar" gorm:"type:text"`
	Terkait     string    `json:"terkait" gorm:"type:text"`
	Slug        string    `json:"slug" gorm:"type:text"`
	TimeInput   time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit    time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}

type PeraturanResult struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Creator     string  `json:"creator"`
	Category    string  `json:"category"`
	Bidang      string  `json:"bidang"`
	SubBidang   string  `json:"sub_bidang"`
	Rank        float64 `json:"rank"`
}

func (a *Peraturan) BeforeCreate(tx *gorm.DB) (err error) {
	a.Slug = a.Category + " " + a.Bidang + " " + a.SubBidang + " " + a.Creator
	return
}

func (a *Peraturan) BeforeUpdate(tx *gorm.DB) (err error) {
	a.Slug = a.Category + " " + a.Bidang + " " + a.SubBidang + " " + a.Creator
	return
}

func (a *Peraturan) Create(data *Peraturan) error {
	//data.Slug = data.Category + " " + data.Bidang + " " + data.SubBidang + " " + data.Creator
	query := db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}

	return nil
}

func (a *Peraturan) Update(ID int, data interface{}) error {
	var err error

	q := db.Model(&Peraturan{}).Where("peraturan_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (a *Peraturan) Delete(ID int) error {
	var err error
	q := db.Where("peraturan_id=?", ID).Delete(&Peraturan{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Peraturan) GetByID(ID int) (*Peraturan, error) {
	var result = &Peraturan{}
	query := db.Where("peraturan_id=?", ID).Find(&result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *Peraturan) TextSearch(term string) (result *[]PeraturanResult, err error) {
	searchTerm := utils.FormatSearch(term)
	queryString := fmt.Sprintf(`select title, description, creator,category, bidang, sub_bidang, ts_rank(text_search, to_tsquery('indonesian','%s')) as rank
	from Peraturan
	where text_search @@ to_tsquery('indonesian','%s')
	order by rank desc`, searchTerm, searchTerm)
	query := db.Raw(queryString).Scan(&result)
	err = query.Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
