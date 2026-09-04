package biz

import (
	"context"
	"log/slog"
	"time"
)

type Article struct {
	Id        int64
	Title     string
	Content   string
	CreatedAT time.Time
	UpdatedAT time.Time
	Like      int64
}

type ArticleRepo interface {
	CreateArticle(context.Context, *Article) (*Article, error)
	UpdateArticle(context.Context, int64, *Article) error
	DeleteArticle(context.Context, int64) error
	GetArticle(context.Context, int64) (*Article, error)
	ListArticle(context.Context) ([]*Article, error)

	//redis
	GetArticleLike(context.Context, int64) (int64, error)
	IncArticleLike(context.Context, int64) error
}

// 创建用例
type ArticleUsecase struct {
	repo   ArticleRepo
	logger *slog.Logger
}

func NewArticleUsecase(repo ArticleRepo, logger *slog.Logger) *ArticleUsecase {
	return &ArticleUsecase{
		repo:   repo,
		logger: logger,
	}
}

func (uc *ArticleUsecase) Create(ctx context.Context, article *Article) (ar *Article, err error) {
	ar, err = uc.repo.CreateArticle(ctx, article)
	if err != nil {
		uc.logger.Error("创建文章失败", "err", err, "path", article.Id)
		return
	}
	return
}

func (uc *ArticleUsecase) Update(ctx context.Context, id int64, article *Article) (err error) {
	return uc.repo.UpdateArticle(ctx, id, article)
}

func (uc *ArticleUsecase) Delete(ctx context.Context, id int64) (err error) {
	return uc.repo.DeleteArticle(ctx, id)
}

func (uc *ArticleUsecase) List(ctx context.Context) (articles []*Article, err error) {
	return uc.repo.ListArticle(ctx)
}

func (uc *ArticleUsecase) Get(ctx context.Context, id int64) (article *Article, err error) {
	article, err = uc.repo.GetArticle(ctx, id)
	if err != nil {
		uc.logger.Error("查询文章失败", "err", err, "path", id)
		return
	}

	// 增加喜欢数量
	err = uc.repo.IncArticleLike(ctx, id)
	if err != nil {
		uc.logger.Error("增加喜欢失败", "err", err, "path", id)
		return
	}

	likeCount, err := uc.repo.GetArticleLike(ctx, id)
	if err != nil {
		uc.logger.Error("获取喜欢失败", "err", err, "path", id)
		return
	}
	article.Like = likeCount
	return
}
