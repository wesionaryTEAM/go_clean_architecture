package infrastructure

import (
	"clean-architecture/lib"
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

// NewBucketStorage creates a new storage client
func NewBucketStorage(logger lib.Logger, env *lib.Env) *storage.Client {
	bucketName := env.StorageBucketName
	ctx := context.Background()
	if bucketName == "" {
		logger.Error("Please check your env file for StorageBucketName")
	}
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("serviceAccountKey.json"))
	if err != nil {
		logger.Fatal(err.Error())
	}
	_, err = client.Bucket(bucketName).Attrs(ctx)
	if err == storage.ErrBucketNotExist {
		logger.Fatalf("Provided bucket %v doesn't exists", bucketName)
	}
	if err != nil {
		logger.Fatalf("Cloud bucket error: %v", err.Error())
	}
	return client
}
