package data

import "time"

type Firm struct {
	FirmID   int       `json:"firm_id" gorm:"primary_key;auto_increment:true"`
	FirmName string    `json:"firm_name" gorm:"type:varchar(100)"`
	Address  string    `json:"address" gorm:"type:varchar(250)"`
	Province string    `json:"province" gorm:"type:varchar(250)"`
	City     string    `json:"city" gorm:"type:varchar(250)"`
	Since    time.Time `json:"time_edit" gorm:"type:timestamp(0)"`
	Lat      float64   `json:"lat" gorm:"type:decimal(10,8)"`
	Lng      float64   `json:"lng" gorm:"type:decimal(11,8)"`
}

type FirmResult struct {
	FirmID   int     `json:"firm_id"`
	FirmName string  `json:"firm_name"`
	Address  string  `json:"address"`
	Province string  `json:"province"`
	City     string  `json:"city"`
	Distance float64 `json:"distance"`
}

func (a *Firm) Create(data *Firm) error {
	//data.Slug = data.Category + " " + data.Bidang + " " + data.SubBidang
	query := db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}

	return nil
}

func (a *Firm) Update(ID int, data interface{}) error {
	var err error

	q := db.Model(&Firm{}).Where("firm_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (a *Firm) Delete(ID int) error {
	var err error
	q := db.Where("firm_id=?", ID).Delete(&Firm{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Firm) GetByID(ID int) (*Firm, error) {
	var result = &Firm{}
	query := db.Where("firm_id=?", ID).Find(&result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *Firm) FindByRadius(lat, lng float64) (result *[]FirmResult, err error) {
	query := db.Raw("SELECT firm_id,firm_name,address,province,city,distance FROM order_by_distance(?,?)", lat, lng).Scan(&result)
	err = query.Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
