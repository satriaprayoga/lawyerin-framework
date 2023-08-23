package firms

import (
	"context"

	"github.com/satriaprayoga/lawyerin-framework/data"
	"github.com/satriaprayoga/lawyerin-framework/pkg/web"
)

type FirmService interface {
	Create(ctx context.Context, data *data.Firm) error
	Update(ctx context.Context, ID int, data interface{}) error
	Delete(ctx context.Context, ID int) error
	GetByID(ctx context.Context, ID int) (*data.Firm, error)
	FindByRadius(ctx context.Context, lat, lng float64) (result web.ResponseModelList, err error)
}
