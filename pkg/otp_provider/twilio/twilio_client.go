package twilio_client

import (
	"example.com/boiletplate/config"
	"github.com/twilio/twilio-go"
)

type TwilioClient struct {
	Twilio           *twilio.RestClient
	VerifyServiceSid string
}

func NewTwilioClient(config *config.Config) *TwilioClient {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.Twilio.TwilioAccountSid,
		Password: config.Twilio.TwilioAuthToken,
	})

	twClient := TwilioClient{client, config.Twilio.VerifyServiceSid}
	return &twClient
}
