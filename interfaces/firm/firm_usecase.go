package firms

import (
	"context"

	"github.com/satriaprayoga/lawyerin-framework/data"
)

type FirmService interface {
	Create(ctx context.Context, data *data.Firm) error
	Update(ctx context.Context, ID int, data interface{}) error
	Delete(ctx context.Context, ID int) error
	GetByID(ctx context.Context, ID int) (*data.Firm, error)
}
