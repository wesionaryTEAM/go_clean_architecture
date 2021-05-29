package services

import (
	"clean-architecture/infrastructure"
	"context"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
)

// BucketService -> handles the file upload/download functions
type BucketService struct {
	logger infrastructure.Logger
	client *storage.Client
	env    infrastructure.Env
}

// NewBucketService -> initilization for the BucketService struct
func NewBucketService(
	logger infrastructure.Logger,
	client *storage.Client,
	env infrastructure.Env,
) BucketService {
	return BucketService{
		logger: logger,
		client: client,
		env:    env,
	}
}

// UploadFile -> uploads the file to the cloud storage
func (s BucketService) UploadFile(
	ctx context.Context,
	file multipart.File,
	fileName string,
	originalFileName string,
) (string, error) {
	var bucketName = s.env.StorageBucketName

	if bucketName == "" {
		s.logger.Fatal("please check your env file for STORAGE_BUCKET_NAME")
	}

	_, err := s.client.Bucket(bucketName).Attrs(ctx)

	if err == storage.ErrBucketNotExist {
		s.logger.Fatalf("provided bucket %v doesn't exists", bucketName)
	}

	if err != nil {
		s.logger.Fatalf("cloud bucket error: %v", err.Error())
	}

	wc := s.client.Bucket(bucketName).Object(fileName).NewWriter(ctx)
	wc.ContentType = "application/octet-stream"
	wc.ContentDisposition = "attachment; filename=" + originalFileName

	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	u, err := url.ParseRequestURI("/" + bucketName + "/" + wc.Attrs().Name)

	if err != nil {
		return "", err
	}

	path := u.EscapedPath()
	path = strings.Replace(path, "/"+bucketName, "", 1)
	path = strings.Replace(path, "/", "", 1)

	return path, nil
}

// GetObjectSignedURL -> gets the signed url for the stored object
func (s BucketService) GetObjectSignedURL(
	object string,
) (string, error) {
	var bucketName = s.env.StorageBucketName

	jsonKey, err := ioutil.ReadFile("serviceAccountKey.json")
	if err != nil {
		return "", nil
	}

	conf, err := google.JWTConfigFromJSON(jsonKey)

	if err != nil {
		return "", err
	}

	opts := &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		Method:         "GET",
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        time.Now().Add(15 * time.Minute),
	}

	u, err := storage.SignedURL(bucketName, object, opts)

	if err != nil {
		return "", err
	}

	return u, nil
}

// RemoveObject -> removes the file from the storage bucket
func (s BucketService) RemoveObject(
	objectName string,
	bucketName string,
) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	objectToDelete := s.client.Bucket(bucketName).Object(objectName)

	if err := objectToDelete.Delete(ctx); err != nil {
		return err
	}

	return nil
}
