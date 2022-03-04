package services

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

type FirebaseService struct {
	*auth.Client
}

// NewFirebaseService : get injected firebase
func NewFirebaseService(fb *auth.Client) FirebaseService {
	return FirebaseService{
		fb,
	}
}

func (fb *FirebaseService) VerifyToken(idToken string) (*auth.Token, error) {
	token, err := fb.VerifyIDToken(context.Background(), idToken)
	return token, err
}

// claims format : gin.H{"user": true}
func (fb *FirebaseService) CreateUserWithClaims(email, password string, claims gin.H) (string, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password)
	u, err := fb.CreateUser(context.Background(), params)
	if err != nil {
		return "", err
	}

	err = fb.SetCustomUserClaims(context.Background(), u.UID, claims)
	return u.UID, err
}

func (fb *FirebaseService) DeleteUserFromFirebase(uid string) error {
	err := fb.DeleteUser(context.Background(), uid)
	return err
}

func (fb *FirebaseService) RetrieveUserByEmail(email string) (*auth.UserRecord, error) {
	user, err := fb.GetUserByEmail(context.Background(), email)
	return user, err
}
