package infrastructure

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewRouter),
	fx.Provide(NewDatabase),
	fx.Provide(NewFBApp),
	fx.Provide(NewFBAuth),
	fx.Provide(NewFirestoreClient),
	fx.Provide(NewBucketStorage),
	fx.Provide(NewFCMClient),
	fx.Provide(NewMigrations),
	fx.Provide(NewGmailService),
)
