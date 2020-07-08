package service

import (
	"context"

	"firebase.google.com/go/auth"
)

type firebaseService struct {
	Firebase *auth.Client
}

// FirebaseService : represent the firebase's services
type FirebaseService interface {
	VerifyToken(idToken string) (*auth.Token, error)
	CreateUser(email, password string) (string, error)
	DeleteUser(uid string) error
}

// NewFirebaseService : get injected firebase
func NewFirebaseService(fb *auth.Client) FirebaseService {
	return &firebaseService{
		Firebase: fb,
	}
}

func (fb *firebaseService) VerifyToken(idToken string) (*auth.Token, error) {
	token, err := fb.Firebase.VerifyIDToken(context.Background(), idToken)
	return token, err
}

func (fb *firebaseService) CreateUser(email, password string) (string, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password)
	u, err := fb.Firebase.CreateUser(context.Background(), params)
	if err != nil {
		return "", err
	}

	claims := map[string]interface{}{"user": true}
	err = fb.Firebase.SetCustomUserClaims(context.Background(), u.UID, claims)
	return u.UID, err
}

func (fb *firebaseService) DeleteUser(uid string) error {
	err := fb.Firebase.DeleteUser(context.Background(), uid)
	return err
}
