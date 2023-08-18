package usecases

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/satriaprayoga/lawyerin-framework/data"
	"github.com/satriaprayoga/lawyerin-framework/interfaces/fileupload"
	"github.com/satriaprayoga/lawyerin-framework/pkg/logger"
)

type FileUploadUsecase struct {
	store   *data.Store
	timeout time.Duration
}

func NewFileUploadUsecase(s *data.Store) fileupload.FileUploadService {
	duration, _ := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	return &FileUploadUsecase{store: s, timeout: time.Duration(duration) * time.Second}
}

func (a *FileUploadUsecase) Create(ctx context.Context, fileUploadData *data.FileUpload) error {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()

	err := a.store.FileUpload.Create(fileUploadData)
	if err != nil {
		log.Error("%v", err)
		return err
	}

	return nil
}

func (a *FileUploadUsecase) Delete(ctx context.Context, ID int) error {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()

	err := a.store.FileUpload.Delete(ID)
	if err != nil {
		log.Error("%v", err)
		return err
	}

	return nil
}

func (a *FileUploadUsecase) GetByID(ctx context.Context, ID int) (*data.FileUpload, error) {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()
	log := logger.GetLogger()
	result, err := a.store.FileUpload.GetByID(ID)
	if err != nil {
		log.Error("%v", err)
		return nil, err
	}
	return result, nil
}
