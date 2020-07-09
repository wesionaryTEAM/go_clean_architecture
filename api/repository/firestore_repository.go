package repository

// import (
// 	"prototype2/entity"
// 	"log"
// 	"context"

// 	"cloud.google.com/go/firestore"
// )

// type repo struct {}

// //New repository constructor
// func NewFirestoreRepository() PostRepository {
// 	return &repo{}
// }

// const (
// 	projectId string = "gocleanarch"
// 	collectionName string = "posts"
// )

// func (*repo) Save(post *entity.Post) (*entity.Post, error) {
// 	c := context.Background()
// 	client, err := firestore.NewClient(c, projectId)
// 	if err != nil {
// 		log.Fatalf("Failed to create a firestore client: %v", err)
// 		return nil, err
// 	}

// 	defer client.Close()
// 	_, _, err = client.Collection(collectionName).Add(c, map[string]interface{}{
// 		"ID": post.ID,
// 		"Title": post.Title,
// 		"Text": post.Text,
// 	})

// 	if err != nil {
// 		log.Fatalf("Failed to create a firestore client: %v", err)
// 		return nil, err
// 	}
// 	return post, nil
// }

// func (*repo) FindAll() ([]entity.Post, error){
// 	c := context.Background()
// 	client, err := firestore.NewClient(c, projectId)
// 	if err != nil {
// 		log.Fatalf("Failed to create a firestore client: %v", err)
// 		return nil, err
// 	}

// 	defer client.Close()
// 	var posts []entity.Post
// 	iterator := client.Collection(collectionName).Documents(c)
// 	for{
// 		doc, err := iterator.Next()
// 		if err != nil {
// 			log.Fatalf("Failed to iterate the list of posts %v", err)
// 			return nil, err
// 		}

// 		post := entity.Post{
// 			ID: doc.Data()["ID"].(int64),
// 			Title: doc.Data()["Title"].(string),
// 			Text: doc.Data()["Text"].(string),
// 		}

// 		posts = append(posts, post)
// 	}
// 	return posts, nil
// }
