package infrastructure

import (
	"clean-architecture/pkg/framework"
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

// NewAWSConfig create a new aws config
func NewAWSConfig(
	env *framework.Env,
) aws.Config {
	c := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
		env.AWSAccessKey, env.AWSSecretAccessKey, ""),
	)
	conf, _ := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(env.AWSRegion),
		config.WithCredentialsProvider(c),
		config.WithClientLogMode(aws.LogRetries),
	)

	return conf
}

func NewCognitoClient(cfg aws.Config) *cognitoidentityprovider.Client {
	return cognitoidentityprovider.NewFromConfig(cfg)
}

func NewSESClient(cfg aws.Config) *sesv2.Client {
	return sesv2.NewFromConfig(cfg)
}

func NewS3Uploader(client *s3.Client) *manager.Uploader {
	return manager.NewUploader(client)
}

// NewPresignClient new presign client
func NewPresignClient(client *s3.Client) *s3.PresignClient {
	return s3.NewPresignClient(client, s3.WithPresignExpires(15*time.Minute))
}

func NewS3Client(cfg aws.Config) *s3.Client {
	return s3.NewFromConfig(cfg)
}
