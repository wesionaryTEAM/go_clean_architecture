package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

type SESService struct {
	*sesv2.Client
}

func NewSESService(client *sesv2.Client) SESService {
	return SESService{
		Client: client,
	}
}

type EmailParams struct {
	From    string
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

func (s SESService) SendEmail(params *EmailParams) error {
	charset := "UTF-8"

	input := &sesv2.SendEmailInput{
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Text: &types.Content{
						Charset: &charset,
						Data:    &params.Body,
					},
				},
				Subject: &types.Content{
					Charset: &charset,
					Data:    &params.Subject,
				},
			},
		},
		Destination: &types.Destination{
			ToAddresses: params.To,
		},
		FromEmailAddress: &params.From,
	}
	if _, err := s.Client.SendEmail(context.Background(), input); err != nil {
		return err
	}
	return nil
}
