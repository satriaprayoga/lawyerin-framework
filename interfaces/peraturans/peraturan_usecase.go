package peraturans

import (
	"context"

	"github.com/satriaprayoga/lawyerin-framework/data"
	"github.com/satriaprayoga/lawyerin-framework/pkg/web"
)

type PeraturanService interface {
	Create(ctx context.Context, data *data.Peraturan) error
	Update(ctx context.Context, ID int, data interface{}) error
	Delete(ctx context.Context, ID int) error
	GetByID(ctx context.Context, ID int) (*data.Peraturan, error)
	TextSearch(ctx context.Context, term string) (result web.ResponseModelList, err error)
}
