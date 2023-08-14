package usecases

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/satriaprayoga/lawyerin-framework/data"
	"github.com/satriaprayoga/lawyerin-framework/interfaces/articles"
	"github.com/satriaprayoga/lawyerin-framework/pkg/logger"
)

type ArticleUsecase struct {
	store   *data.Store
	timeout time.Duration
}

func NewArticleUsecase(s *data.Store) articles.ArticleService {
	duration, _ := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	return &ArticleUsecase{store: s, timeout: time.Duration(duration) * time.Second}
}

func (a *ArticleUsecase) Create(ctx context.Context, data *data.Article) error {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()
	err := a.store.Article.Create(data)
	if err != nil {
		log.Error("%v", err)
		return err
	}

	return nil
}
