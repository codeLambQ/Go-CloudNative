package service

import (
	"context"
	"moddy-blog-article/api/article/v1"
	"moddy-blog-article/internal/biz"
)

// GreeterService is a greeter service.
type ArticleService struct {
}

func (a ArticleService) CreateArticle(ctx context.Context, request *v1.CreateArticleRequest) (*v1.CreateArticleReply, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleService) DeleteArticle(ctx context.Context, request *v1.DeleteArticleRequest) (*v1.DeleteArticleReply, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleService) GetArticle(ctx context.Context, request *v1.GetArticleRequest) (*v1.GetArticleReply, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleService) ListArticle(ctx context.Context, request *v1.ListArticleRequest) (*v1.ListArticleReply, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleService) UpdateArticle(ctx context.Context, request *v1.UpdateArticleRequest) (*v1.UpdateArticleReply, error) {
	//TODO implement me
	panic("implement me")
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.ArticleUsecase) *ArticleService {
	return &ArticleService{}
}
