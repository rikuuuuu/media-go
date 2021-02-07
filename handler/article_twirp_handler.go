package handler

import (
	"context"
	"fmt"
	pb "mediago/pb"
	"mediago/repository"
)

type ArticleService struct {
	contextProvider ContextProvider
}

type ArticleHandler interface {
	Get(ctx context.Context, req *pb.ArticleID) (resp *pb.Article, err error)
	GetAll(ctx context.Context, req *pb.Pager) (resp *pb.ArticleList, err error)
	Create(ctx context.Context, req *pb.CreateArticleRequest) (resp *pb.Article, err error)
	Update(ctx context.Context, req *pb.UpdateArticleRequest) (resp *pb.Article, err error)
	Delete(ctx context.Context, req *pb.ArticleID) (resp *pb.Empty, err error)
}

func NewArticleService() AdminUserHandler {
	c := NewContextProvider()
	return &AdminUserService{contextProvider: c}
}

func (a *ArticleService) Get(ctx context.Context, req *pb.ArticleID) (resp *pb.Article, err error) {

	fmt.Printf("%v", "article get")

	article, err := repository.ArticleGet(ctx, req)

	return article, err
}

func (a *ArticleService) GetAll(ctx context.Context, req *pb.Pager) (resp *pb.ArticleList, err error) {

	fmt.Printf("%v", "article getall")

	articles, err := repository.ArticleGetAll(ctx, req)

	return articles, err
}

func (a *ArticleService) Create(ctx context.Context, req *pb.CreateArticleRequest) (resp *pb.Article, err error) {

	fmt.Printf("%v", "article create")

	article, err := repository.ArticleCreate(ctx, req)

	return article, err
}

func (a *ArticleService) Update(ctx context.Context, req *pb.UpdateArticleRequest) (resp *pb.Article, err error) {

	fmt.Printf("%v", "article update")

	article, err := repository.ArticleUpdate(ctx, req)

	return article, err
}

func (a *ArticleService) Delete(ctx context.Context, req *pb.ArticleID) (resp *pb.Empty, err error) {

	fmt.Printf("%v", "article delete")

	empty, err := repository.ArticleDelete(ctx, req)

	return empty, err
}