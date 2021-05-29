package infrastructure

import (
	"context"
	"path/filepath"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

// NewFBApp creates new firebase app instance
func NewFBApp(logger Logger) *firebase.App {

	ctx := context.Background()

	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		logger.Panic("Unable to load serviceAccountKey.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		logger.Fatalf("Firebase NewApp: %v", err)
	}
	logger.Info("✅ Firebase app initialized.")
	return app
}

// NewFBAuth creates new firebase auth client
func NewFBAuth(logger Logger, app *firebase.App) *auth.Client {

	ctx := context.Background()

	firebaseAuth, err := app.Auth(ctx)
	if err != nil {
		logger.Fatalf("Firebase Authentication: %v", err)
	}

	return firebaseAuth
}

// NewFirestoreClient creates new firestore client
func NewFirestoreClient(logger Logger, app *firebase.App) *firestore.Client {
	ctx := context.Background()

	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		logger.Fatalf("Firestore client: %v", err)
	}

	return firestoreClient
}

// NewFCMClient creates new firebase cloud messaging client
func NewFCMClient(logger Logger, app *firebase.App) *messaging.Client {
	ctx := context.Background()
	messagingClient, err := app.Messaging(ctx)
	if err != nil {
		logger.Fatalf("Firebase messaing: %v", err)
	}
	return messagingClient
}
