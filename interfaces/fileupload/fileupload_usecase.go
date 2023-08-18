package fileupload

import (
	"context"

	"github.com/satriaprayoga/lawyerin-framework/data"
)

type FileUploadService interface {
	Create(ctx context.Context, data *data.FileUpload) error
	Delete(ctx context.Context, ID int) error
	GetByID(ctx context.Context, ID int) (*data.FileUpload, error)
}
