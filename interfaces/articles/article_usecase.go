package articles

import (
	"context"

	"github.com/satriaprayoga/lawyerin-framework/data"
)

type ArticleService interface {
	Create(ctx context.Context, data *data.Article) error
}
