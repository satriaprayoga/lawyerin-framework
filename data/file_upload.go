package data

import "time"

type FileUpload struct {
	FileID    int       `json:"file_id" gorm:"primary_key;auto_increment:true"`
	FileName  string    `json:"file_name" gorm:"type:varchar(60)"`
	FilePath  string    `json:"file_path" gorm:"type:varchar(150)"`
	FileType  string    `json:"file_type" gorm:"type:varchar(10)"`
	UserInput string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit  string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit  time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}

func (a *FileUpload) Create(data *FileUpload) error {
	query := db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}

	return nil
}

func (a *FileUpload) Update(ID int, data interface{}) error {
	var err error

	q := db.Model(&FileUpload{}).Where("file_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (a *FileUpload) Delete(ID int) error {
	var err error
	q := db.Where("file_id=?", ID).Delete(&FileUpload{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}

func (a *FileUpload) GetByID(ID int) (*FileUpload, error) {
	var result = &FileUpload{}
	query := db.Where("file_id=?", ID).Find(&result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
