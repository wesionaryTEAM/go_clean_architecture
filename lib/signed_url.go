package lib

import (
	"encoding/json"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
)

type SignedURL string

// UnmarshalJSON -> convert from json string
func (s *SignedURL) UnmarshalJSON(by []byte) error {
	str := ""
	_ = json.Unmarshal(by, &str)
	*s = SignedURL(str)
	return nil
}

// MarshalJSON -> convert to json string
func (s SignedURL) MarshalJSON() ([]byte, error) {
	signedURL, err := s.getObjectSignedURL()
	if err != nil {
		return []byte("\"\""), nil
	}

	str := "\"" + signedURL + "\""
	return []byte(str), nil
}

// GetObjectSignedURL -> gets the signed url for the stored object
func (s SignedURL) getObjectSignedURL() (string, error) {
	env := GetEnv()
	bucketName := env.StorageBucketName

	jsonKey, err := os.ReadFile("serviceAccountKey.json")
	if err != nil {
		return "", nil
	}

	conf, err := google.JWTConfigFromJSON(jsonKey)

	if err != nil {
		return "", err
	}

	opts := &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		Method:         "GET",
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        time.Now().Add(15 * time.Minute),
	}
	if s == "" {
		return "", nil
	}
	u, err := storage.SignedURL(bucketName, string(s), opts)

	if err != nil {
		return "", err
	}

	return u, nil
}
