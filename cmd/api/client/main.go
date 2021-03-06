package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	article "github.com/tanimutomo/grpcapi-go-server/pkg/grpcs/article"
)

const host = "localhost:50051"

func main() {
	doArticleRequests()
}

func doArticleRequests() {
	fmt.Println("do articles")

	conn, err := grpc.Dial(
		host,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("error in connection")
		return
	}
	defer conn.Close()
	c := article.NewArticleServiceClient(conn)

	if err := articleGetArticleRequest(c, uint64(1)); err != nil {
		log.Fatalf("error in Get: %v\n", err)
		return
	}
	if err := articleListArticlesRequest(c); err != nil {
		log.Fatalf("error in List: %v\n", err)
		return
	}
	if err := articleCreateArticleRequest(c, "title"); err != nil {
		log.Fatalf("error in Create: %v\n", err)
		return
	}

	fmt.Println("\nend articles")
}

func articleGetArticleRequest(client article.ArticleServiceClient, id uint64) error {
	fmt.Println("\ndo article/Get")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	req := article.GetArticleRequest{
		Id: id,
	}
	res, err := client.GetArticle(ctx, &req)
	if err != nil {
		return errors.Wrap(err, "failed to receive response")
	}
	log.Printf("response: %+v\n", res.GetArticle())
	return nil
}

func articleListArticlesRequest(client article.ArticleServiceClient) error {
	fmt.Println("\ndo article/List")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	req := article.ListArticlesRequest{}
	res, err := client.ListArticles(ctx, &req)
	if err != nil {
		return errors.Wrap(err, "failed to receive response")
	}
	log.Printf("response: %+v\n", res.GetArticles())
	return nil
}

func articleCreateArticleRequest(client article.ArticleServiceClient, title string) error {
	fmt.Println("\ndo article/Create")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	req := article.CreateArticleRequest{
		Title: title,
	}
	res, err := client.CreateArticle(ctx, &req)
	if err != nil {
		return errors.Wrap(err, "failed to receive response")
	}
	log.Printf("response: %+v\n", res.GetArticle())
	return nil
}
