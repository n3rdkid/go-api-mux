package repository

import (
	"context"
	"go-api-mux/entity"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// PostRepository ->
type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

// NewPostRepository ->
func NewPostRepository() PostRepository {
	return &repo{}
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	log.Printf("CLIENT %v", client.Collection(os.Getenv("COLLECTION_NAME")))
	if err != nil {
		log.Fatalf("Failed to create a firestore client : %v", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(os.Getenv("COLLECTION_NAME")).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		log.Fatalf("Failed adding a new post : %v", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("Failed to create a firestore client : %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post
	itr := client.Collection(os.Getenv("COLLECTION_NAME")).Documents(ctx)
	for {
		document, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate posts collection : %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    document.Data()["ID"].(int64),
			Title: document.Data()["Title"].(string),
			Text:  document.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
