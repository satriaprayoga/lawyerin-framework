package usecases

import (
	"context"

	"github.com/satriaprayoga/lawyerin-framework/data"
	"github.com/satriaprayoga/lawyerin-framework/interfaces/articles"
)

type ArticleUsecase struct {
	store *data.Store
}

func NewArticleUsecase(s *data.Store) articles.ArticleService {
	return &ArticleUsecase{store: s}
}

func (a *ArticleUsecase) Create(ctx context.Context, data *data.Book) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	err := a.store.Book.Create(data)
	if err != nil {
		return err
	}

	return nil
}
