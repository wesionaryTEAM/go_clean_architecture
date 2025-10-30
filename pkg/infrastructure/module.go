package infrastructure

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(
		NewRouter,
		NewDBProvider,
		NewDatabase,
		//NewS3Client,
		//NewAWSConfig,
		//NewPresignClient,
		//NewS3Uploader,
		//NewCognitoClient,
		//NewSESClient,
	),
)
