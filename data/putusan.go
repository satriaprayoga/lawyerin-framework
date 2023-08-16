package data

import (
	"fmt"
	"time"

	"github.com/satriaprayoga/lawyerin-framework/pkg/utils"
)

type Putusan struct {
	PutusanID   int    `json:"putusan_id" gorm:"primary_key;auto_increment:true"`
	Title       string `json:"title" gorm:"type:text"`
	Description string `json:"description" gorm:"type:text"`
	Category    string `json:"category" gorm:"type:varchar(100)"`
	Bidang      string `json:"bidang" gorm:"type:varchar(100)"`
	SubBidang   string `json:"sub_bidang" gorm:"type:varchar(100)"`
	Slug        string `json:"slug" gorm:"type:text"`

	Jenis      string    `json:"jenis" gorm:"type:varchar(100)"`
	Tingkat    string    `json:"tingkat" gorm:"type:varchar(100)"`
	DatePts    time.Time `json:"date_pts" gorm:"type:timestamp(0)"`
	Ketua      string    `json:"ketua" gorm:"type:varchar(200)"`
	Anggota    string    `json:"anggota" gorm:"type:varchar(200)"`
	Amar       string    `json:"amar" gorm:"type:varchar(100)"`
	AmarDetail string    `json:"amar_details" gorm:"type:text"`
	Pemohon    string    `json:"pemohon" gorm:"type:varchar(200)"`
	Vs         string    `json:"versus" gorm:"type:varchar(200)"`
	UserInput  string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit   string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput  time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit   time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}

type PutusanResult struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Creator     string  `json:"creator"`
	Category    string  `json:"category"`
	Bidang      string  `json:"bidang"`
	SubBidang   string  `json:"sub_bidang"`
	Rank        float64 `json:"rank"`
}

func (a *Putusan) Create(data *Putusan) error {
	data.Slug = data.Category + " " + data.Bidang + " " + data.SubBidang + " " + data.Pemohon + " " + data.Vs
	query := db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}

	return nil
}

func (a *Putusan) Update(ID int, data interface{}) error {
	var err error

	q := db.Model(&Putusan{}).Where("putusan_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (a *Putusan) Delete(ID int) error {
	var err error
	q := db.Where("putusan_id=?", ID).Delete(&Putusan{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Putusan) GetByID(ID int) (*Putusan, error) {
	var result = &Putusan{}
	query := db.Where("putusan_id=?", ID).Find(&result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *Putusan) TextSearch(term string) (result *[]PutusanResult, err error) {
	searchTerm := utils.FormatSearch(term)
	queryString := fmt.Sprintf(`select title, description, creator,category, bidang, sub_bidang, ts_rank(text_search, to_tsquery('indonesian','%s')) as rank
	from putusan
	where text_search @@ to_tsquery('indonesian','%s')
	order by rank desc`, searchTerm, searchTerm)
	query := db.Raw(queryString).Scan(&result)
	err = query.Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
