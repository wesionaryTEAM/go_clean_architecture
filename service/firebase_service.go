package service

import (
	"context"

	"firebase.google.com/go/auth"
)

type firebaseService struct {
	FB *auth.Client
}

// FirebaseService : represent the firebase's services
type FirebaseService interface {
	VerifyToken(idToken string) (*auth.Token, error)
	CreateUser(email, password string) error
}

// NewFirebaseService : get injected firebase
func NewFirebaseService(fb *auth.Client) FirebaseService {
	return &firebaseService{
		FB: fb,
	}
}

func (f *firebaseService) VerifyToken(idToken string) (*auth.Token, error) {
	token, err := f.FB.VerifyIDToken(context.Background(), idToken)
	return token, err
}

func (f *firebaseService) CreateUser(email, password string) error {
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password)
	_, err := f.FB.CreateUser(context.Background(), params)
	return err
}
