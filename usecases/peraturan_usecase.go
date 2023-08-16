package usecases

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/satriaprayoga/lawyerin-framework/data"
	"github.com/satriaprayoga/lawyerin-framework/interfaces/peraturans"
	"github.com/satriaprayoga/lawyerin-framework/pkg/logger"
	"github.com/satriaprayoga/lawyerin-framework/pkg/web"
)

type PeraturanUsecase struct {
	store   *data.Store
	timeout time.Duration
}

func NewPeraturanUsecase(s *data.Store) peraturans.PeraturanService {
	duration, _ := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	return &PeraturanUsecase{store: s, timeout: time.Duration(duration) * time.Second}
}

func (a *PeraturanUsecase) Create(ctx context.Context, peraturanData *data.Peraturan) error {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()

	err := a.store.Peraturan.Create(peraturanData)
	if err != nil {
		log.Error("%v", err)
		return err
	}

	return nil
}

func (a *PeraturanUsecase) Update(ctx context.Context, ID int, data interface{}) error {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()

	err := a.store.Peraturan.Update(ID, data)
	if err != nil {
		log.Error("%v", err)
		return err
	}

	return nil
}

func (a *PeraturanUsecase) Delete(ctx context.Context, ID int) error {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()

	err := a.store.Peraturan.Delete(ID)
	if err != nil {
		log.Error("%v", err)
		return err
	}

	return nil
}

func (a *PeraturanUsecase) GetByID(ctx context.Context, ID int) (*data.Peraturan, error) {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()
	log := logger.GetLogger()
	result, err := a.store.Peraturan.GetByID(ID)
	if err != nil {
		log.Error("%v", err)
		return nil, err
	}
	return result, nil
}

func (a *PeraturanUsecase) TextSearch(ctx context.Context, term string) (result web.ResponseModelList, err error) {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()
	result.Data, err = a.store.Peraturan.TextSearch(term)
	if err != nil {
		log.Error("%v", err)
		return result, err
	}
	return result, err
}
