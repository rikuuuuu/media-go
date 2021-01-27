package repository

import (
	// "time"
	"context"
	"log"
	"mediago/model"
	"cloud.google.com/go/firestore"
	pb "mediago/pb"
)

func ArticleGet(ctx context.Context, req *pb.ArticleID) (resp *pb.Article, err error) {

	client, err := firebaseInit(ctx)
		if err != nil {
			log.Fatal(err)
	}

	dsnap, err := client.Collection("article").Doc(req.Id).Get(ctx)
	if err != nil {
        return nil, err
	}
	dsnap.DataTo(&resp)

	defer client.Close()

	return &pb.Article{
		Id: resp.Id,
		UserID: resp.UserID,
		Title: resp.Title,
		Description: resp.Description,
		CreatedAt: resp.CreatedAt,
		ThumbnailURL: resp.ThumbnailURL,
	}, nil
}

func ArticleGetAll(ctx context.Context, req *pb.Pager) (resp *pb.ArticleList, err error) {

	client, err := firebaseInit(ctx)
		if err != nil {
			log.Fatal(err)
	}

	page := int(req.GetPage())
	limit := page * 10

	// allData := client.Collection("article")
	docs, err := client.Collection("article").OrderBy("createdAt", firestore.Desc).Limit(limit).Documents(ctx).GetAll()
	if err != nil {
		log.Fatalf("Failed adding getAll: %v", err)
	}

	articles := make([]*pb.Article, 0, len(docs))
	for _, doc := range docs {
		a := new(pb.Article)
		mapToStruct(doc.Data(), &a)
		articles = append(articles, a)
	}

	return &pb.ArticleList{
				Items: articles,
			}, nil
}


func ArticleCreate(ctx context.Context, req *pb.CreateArticleRequest) (resp *pb.Article, err error) {

	// now := time.Now()
  
	// article.Created = now
	// article.Updated = now

	client, err := firebaseInit(ctx)
		if err != nil {
			log.Fatal(err)
	}

	ref := client.Collection("article").NewDoc()

	article := &model.CreateArticle{
		Id: ref.ID,
		UserId: "test",
		Title: req.GetTitle(),
		Description: req.GetDescription(),
		CreatedAt: "test",
		ThumbnailURL: req.GetThumbnailURL(),
	}

	_, err = ref.Set(ctx, article)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	defer client.Close()
  
	return &pb.Article{
		Id: "test",
		UserID: "test",
		Title: req.GetTitle(),
		Description: req.GetDescription(),
		CreatedAt: "test",
		ThumbnailURL: req.GetThumbnailURL(),
	}, nil
}

func ArticleUpdate(ctx context.Context, req *pb.UpdateArticleRequest) (resp *pb.Article, err error) {

	client, err := firebaseInit(ctx)
		if err != nil {
			log.Fatal(err)
	}

	_, err = client.Collection("article").Doc(req.GetId()).Set(ctx, map[string]interface{}{
		"id": req.GetId(),
		"title": req.GetTitle(),
		"description": req.GetDescription(),
		"thumbnailURL": req.GetThumbnailURL(),
	}, firestore.MergeAll)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	dsnap, err := client.Collection("article").Doc(req.GetId()).Get(ctx)
	if err != nil {
        return nil, err
	}
	dsnap.DataTo(&resp)

	defer client.Close()

	return &pb.Article{
		Id: resp.Id,
		UserID: resp.UserID,
		Title: resp.Title,
		Description: resp.Description,
		CreatedAt: resp.CreatedAt,
		ThumbnailURL: resp.ThumbnailURL,
	}, nil
}

func ArticleDelete(ctx context.Context, req *pb.ArticleID) (resp *pb.Empty, err error) {

	client, err := firebaseInit(ctx)
		if err != nil {
			log.Fatal(err)
	}

	_, err = client.Collection("article").Doc(req.GetId()).Delete(ctx)
	if err != nil {
			log.Printf("An error has occurred: %s", err)
	}

	defer client.Close()

	return &pb.Empty{}, nil
}