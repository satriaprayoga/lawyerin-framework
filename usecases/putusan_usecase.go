package usecases

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/satriaprayoga/lawyerin-framework/data"
	"github.com/satriaprayoga/lawyerin-framework/interfaces/putusans"
	"github.com/satriaprayoga/lawyerin-framework/pkg/logger"
	"github.com/satriaprayoga/lawyerin-framework/pkg/web"
)

type PutusanUsecase struct {
	store   *data.Store
	timeout time.Duration
}

func NewPutusanUsecase(s *data.Store) putusans.PutusanService {
	duration, _ := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	return &PutusanUsecase{store: s, timeout: time.Duration(duration) * time.Second}
}

func (a *PutusanUsecase) Create(ctx context.Context, putusanData *data.Putusan) error {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()

	err := a.store.Putusan.Create(putusanData)
	if err != nil {
		log.Error("%v", err)
		return err
	}

	return nil
}

func (a *PutusanUsecase) Update(ctx context.Context, ID int, data interface{}) error {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()

	err := a.store.Putusan.Update(ID, data)
	if err != nil {
		log.Error("%v", err)
		return err
	}

	return nil
}

func (a *PutusanUsecase) Delete(ctx context.Context, ID int) error {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()

	err := a.store.Putusan.Delete(ID)
	if err != nil {
		log.Error("%v", err)
		return err
	}

	return nil
}

func (a *PutusanUsecase) GetByID(ctx context.Context, ID int) (*data.Putusan, error) {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()
	log := logger.GetLogger()
	result, err := a.store.Putusan.GetByID(ID)
	if err != nil {
		log.Error("%v", err)
		return nil, err
	}
	return result, nil
}

func (a *PutusanUsecase) TextSearch(ctx context.Context, term string) (result web.ResponseModelList, err error) {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()
	result.Data, err = a.store.Putusan.TextSearch(term)
	if err != nil {
		log.Error("%v", err)
		return result, err
	}
	return result, err
}
