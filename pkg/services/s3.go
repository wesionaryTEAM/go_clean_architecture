package services

import (
	"clean-architecture/pkg/framework"
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Service struct {
	logger framework.Logger
	env    *framework.Env
	client *s3.Client
}

func NewS3Service(
	logger framework.Logger,
	env *framework.Env,
	client *s3.Client,
) S3Service {
	return S3Service{
		logger: logger,
		env:    env,
		client: client,
	}
}

func (s *S3Service) UploadFile(
	ctx context.Context,
	file io.Reader,
	fileName string,
) (string, error) {

	bucketName := s.env.StorageBucketName

	if bucketName == "" {
		s.logger.Fatal("Bucket name missing.")
	}

	_, err := s.client.HeadBucket(context.Background(), &s3.HeadBucketInput{
		Bucket: &bucketName,
	})

	if err != nil {
		s.logger.Fatalf("%v Bucket doesn't exists. Error because of %v", bucketName, err.Error())
	}

	_, err = s.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &fileName,
		Body:   file,
	})

	if err != nil {
		s.logger.Fatal("Failed to upload the file in the bucket.", err.Error())
		return "", err
	}

	return fileName, nil
}

func (s *S3Service) GetSignedURL(key string) (string, error) {
	bucketName := s.env.StorageBucketName

	if bucketName == "" {
		s.logger.Fatal("Bucket name missing.")
	}

	_, err := s.client.HeadBucket(context.Background(), &s3.HeadBucketInput{
		Bucket: &bucketName,
	})

	if err != nil {
		s.logger.Fatalf("%v Bucket doesn't exists. Error because of %v", bucketName, err.Error())
	}

	presignClient := s3.NewPresignClient(s.client)
	presignedUrl, err := presignClient.PresignGetObject(context.Background(),
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(key),
		},
		s3.WithPresignExpires(time.Minute*15))
	if err != nil {
		s.logger.Fatal(err)
	}
	return presignedUrl.URL, nil

}
