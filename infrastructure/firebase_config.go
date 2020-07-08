package infrastructure

import (
	"context"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

// SetupFirebase , up firebase client
func InitializeFirebase() *auth.Client {
	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		panic("Unable to load serviceAccountKey.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	firebaseAuth, err := app.Auth(context.Background())
	if err != nil {
		panic("Firebase load error")
	}

	return firebaseAuth
}
