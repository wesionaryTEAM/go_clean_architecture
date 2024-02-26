package infrastructure

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(
		NewRouter,
		NewDatabase,
		NewFBApp,
		NewFBAuth,
		NewFirestoreClient,
		NewBucketStorage,
		NewFCMClient,
		NewMigrations,
		NewS3Client,
		NewAWSConfig,
		NewPresignClient,
		NewS3Uploader,
		NewCognitoClient,
	),
)
