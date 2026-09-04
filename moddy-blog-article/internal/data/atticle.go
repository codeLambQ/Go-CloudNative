package data

import (
	"context"
	"fmt"
	"log/slog"
	"moddy-blog-article/internal/biz"
)

type articleRepo struct {
	data *Data
	log  *slog.Logger
}

func NewArticleRepo(data *Data, logger *slog.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  logger,
	}
}

func (ar *articleRepo) CreateArticle(ctx context.Context, article *biz.Article) (*biz.Article, error) {
	art, err := ar.data.PDB.Article.Create().SetTitle(article.Title).SetContent(article.Content).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &biz.Article{
		Id:        art.ID,
		Title:     art.Title,
		Content:   art.Content,
		CreatedAT: art.CreatedAt,
		UpdatedAT: art.UpdatedAt,
	}, err
}
func (ar *articleRepo) UpdateArticle(ctx context.Context, id int64, article *biz.Article) error {
	_, err := ar.data.PDB.Article.UpdateOneID(id).SetTitle(article.Title).SetContent(article.Content).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ar *articleRepo) DeleteArticle(ctx context.Context, id int64) error {
	if err := ar.data.PDB.Article.DeleteOneID(id).Exec(ctx); err != nil {
		ar.log.Error("删除失败")
		return err
	}
	return nil
}
func (ar *articleRepo) GetArticle(ctx context.Context, id int64) (*biz.Article, error) {
	article, err := ar.data.PDB.Article.Get(ctx, id)
	if err != nil {
		ar.log.Error("获取文章失败")
		return nil, err
	}
	likeCount, err := ar.GetArticleLike(ctx, id)
	if err != nil {
		ar.log.Error("获取喜欢失败")
		return nil, err
	}
	return &biz.Article{
		Id:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Like:      likeCount,
		CreatedAT: article.CreatedAt,
		UpdatedAT: article.UpdatedAt,
	}, nil
}
func (ar *articleRepo) ListArticle(ctx context.Context) ([]*biz.Article, error) {
	articles, err := ar.data.PDB.Article.Query().All(ctx)
	if err != nil {
		ar.log.Error("获取文章失败")
		return nil, err
	}
	arts := make([]*biz.Article, 10)
	for _, art := range articles {
		likeCount, _ := ar.GetArticleLike(ctx, art.ID)
		arts = append(arts, &biz.Article{
			Id:        art.ID,
			Title:     art.Title,
			Content:   art.Content,
			Like:      likeCount,
			CreatedAT: art.CreatedAt,
			UpdatedAT: art.UpdatedAt,
		})
	}
	return arts, nil
}

// redis

func likeKey(id int64) string {
	return fmt.Sprintf("like:%d", id)
}
func (ar *articleRepo) GetArticleLike(ctx context.Context, id int64) (int64, error) {
	likeS := ar.data.redisDB.Get(ctx, likeKey(id))
	likeCount, err := likeS.Int64()
	if err != nil {
		ar.log.Error("转换失败")
		return 0, err
	}
	return likeCount, nil
}
func (ar *articleRepo) IncArticleLike(ctx context.Context, id int64) error {
	_, err := ar.data.redisDB.Incr(ctx, likeKey(id)).Result()
	if err != nil {
		ar.log.Error("转换失败")
		return err
	}
	return nil
}
