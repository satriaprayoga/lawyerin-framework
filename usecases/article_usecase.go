package usecases

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/satriaprayoga/lawyerin-framework/data"
	"github.com/satriaprayoga/lawyerin-framework/interfaces/articles"
	"github.com/satriaprayoga/lawyerin-framework/pkg/logger"
	"github.com/satriaprayoga/lawyerin-framework/pkg/web"
)

type ArticleUsecase struct {
	store   *data.Store
	timeout time.Duration
}

func NewArticleUsecase(s *data.Store) articles.ArticleService {
	duration, _ := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	return &ArticleUsecase{store: s, timeout: time.Duration(duration) * time.Second}
}

func (a *ArticleUsecase) Create(ctx context.Context, articleData *data.Article) error {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()

	err := a.store.Article.Create(articleData)
	if err != nil {
		log.Error("%v", err)
		return err
	}

	return nil
}

func (a *ArticleUsecase) Update(ctx context.Context, ID int, data interface{}) error {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()

	err := a.store.Article.Update(ID, data)
	if err != nil {
		log.Error("%v", err)
		return err
	}

	return nil
}

func (a *ArticleUsecase) Delete(ctx context.Context, ID int) error {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()

	err := a.store.Article.Delete(ID)
	if err != nil {
		log.Error("%v", err)
		return err
	}

	return nil
}

func (a *ArticleUsecase) GetByID(ctx context.Context, ID int) (*data.Article, error) {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()
	log := logger.GetLogger()
	result, err := a.store.Article.GetByID(ID)
	if err != nil {
		log.Error("%v", err)
		return nil, err
	}
	return result, nil
}

func (a *ArticleUsecase) TextSearch(ctx context.Context, term string) (result web.ResponseModelList, err error) {
	_, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	log := logger.GetLogger()
	result.Data, err = a.store.Article.TextSearch(term)
	if err != nil {
		log.Error("%v", err)
		return result, err
	}
	return result, err
}
