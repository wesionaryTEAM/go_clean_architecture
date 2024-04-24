package infrastructure

import (
	"clean-architecture/pkg/framework"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// NewAWSConfig create a new aws config
func NewAWSConfig(
	env *framework.Env,
) *aws.Config {
	c := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
		env.AWSAccessKey, env.AWSSecretAccessKey, ""),
	)
	conf, _ := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(env.AWSRegion),
		config.WithCredentialsProvider(c),
		config.WithClientLogMode(aws.LogRetries),
	)

	return &conf
}

func NewCognitoClient(cfg *aws.Config) *cognitoidentityprovider.Client {
	return cognitoidentityprovider.NewFromConfig(*cfg)
}
