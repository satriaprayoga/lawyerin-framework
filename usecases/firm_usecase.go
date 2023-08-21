package usecases

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/satriaprayoga/lawyerin-framework/data"
	firms "github.com/satriaprayoga/lawyerin-framework/interfaces/firm"
	"github.com/satriaprayoga/lawyerin-framework/pkg/logger"
)

type FirmUsecase struct {
	store   *data.Store
	timeout time.Duration
}

func NewFirmUsecase(s *data.Store) firms.FirmService {
	duration, _ := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	return &FirmUsecase{store: s, timeout: time.Duration(duration) * time.Second}
}

func (a *FirmUsecase) Create(ctx context.Context, firmData *data.Firm) error {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()

	err := a.store.Firm.Create(firmData)
	if err != nil {
		log.Error("%v", err)
		return err
	}

	return nil
}

func (a *FirmUsecase) Update(ctx context.Context, ID int, data interface{}) error {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()

	err := a.store.Firm.Update(ID, data)
	if err != nil {
		log.Error("%v", err)
		return err
	}

	return nil
}

func (a *FirmUsecase) Delete(ctx context.Context, ID int) error {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()

	err := a.store.Firm.Delete(ID)
	if err != nil {
		log.Error("%v", err)
		return err
	}

	return nil
}

func (a *FirmUsecase) GetByID(ctx context.Context, ID int) (*data.Firm, error) {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()
	log := logger.GetLogger()
	result, err := a.store.Firm.GetByID(ID)
	if err != nil {
		log.Error("%v", err)
		return nil, err
	}
	return result, nil
}
