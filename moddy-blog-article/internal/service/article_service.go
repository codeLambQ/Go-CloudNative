package service

import (
	"context"
	"log/slog"
	pb "moddy-blog-article/api/article/v1"
	"moddy-blog-article/internal/biz"
)

type ArticleService struct {
	pb.UnsafeArticleServerServer

	articleUseCase *biz.ArticleUsecase
	log            *slog.Logger
}

func NewArticleService(articleUseCase *biz.ArticleUsecase, log *slog.Logger) *ArticleService {
	return &ArticleService{
		articleUseCase: articleUseCase,
		log:            log,
	}
}

func (as *ArticleService) CreateArticle(ctx context.Context, request *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
	as.log.Info("开始添加文章")
	article := &biz.Article{
		Title:   request.Title,
		Content: request.Content,
	}
	_, err := as.articleUseCase.Create(ctx, article)
	if err != nil {
		as.log.Error("添加文章失败")
		return nil, err
	}
	as.log.Info("添加文章成功")
	return &pb.CreateArticleReply{}, err

}

func (as *ArticleService) UpdateArticle(ctx context.Context, request *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	as.log.Info("开始更新文章")
	article := &biz.Article{
		Id:      request.Id,
		Title:   request.Title,
		Content: request.Content,
	}
	err := as.articleUseCase.Update(ctx, request.Id, article)
	if err != nil {
		as.log.Error("更新文章失败")
		return nil, err
	}
	as.log.Info("更新文章成功")
	return &pb.UpdateArticleReply{}, err
}

func (as *ArticleService) DeleteArticle(ctx context.Context, request *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
	as.log.Info("删除文章开始")
	err := as.articleUseCase.Delete(ctx, request.Id)
	if err != nil {
		as.log.Error("删除文章失败")
		return nil, err
	}
	as.log.Info("删除文章成功")

	return &pb.DeleteArticleReply{}, err
}

func (as *ArticleService) GetArticle(ctx context.Context, request *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	as.log.Info("开始获取文章")
	article, err := as.articleUseCase.Get(ctx, request.Id)
	if err != nil {
		as.log.Error("获取文章失败")
		return nil, err
	}
	as.log.Error("获取文章成功")
	return &pb.GetArticleReply{
		Article: &pb.Article{
			Id:      article.Id,
			Title:   article.Title,
			Content: article.Content,
			Like:    article.Like,
		},
	}, err
}

func (as *ArticleService) ListArticle(ctx context.Context, request *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	as.log.Info("开始获取文章列表")
	articles, err := as.articleUseCase.List(ctx)
	if err != nil {
		as.log.Error("获取文章列表失败")
		return nil, err
	}
	reply := &pb.ListArticleReply{}

	for _, article := range articles {
		reply.Articles = append(reply.Articles, &pb.Article{
			Id:      article.Id,
			Title:   article.Title,
			Content: article.Content,
			Like:    article.Like,
		})
	}
	as.log.Error("获取文章列表成功")
	return reply, err
}
