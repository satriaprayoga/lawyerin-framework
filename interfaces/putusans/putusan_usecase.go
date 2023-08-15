package putusans

import (
	"context"

	"github.com/satriaprayoga/lawyerin-framework/data"
	"github.com/satriaprayoga/lawyerin-framework/pkg/web"
)

type PutusanService interface {
	Create(ctx context.Context, data *data.Putusan) error
	Update(ctx context.Context, ID int, data interface{}) error
	Delete(ctx context.Context, ID int) error
	GetByID(ctx context.Context, ID int) (*data.Putusan, error)
	TextSearch(ctx context.Context, term string) (result web.ResponseModelList, err error)
}
