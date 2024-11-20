package otphandler

import (
	"errors"
	"fmt"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type TwilioAdapter struct {
	client           *twilio.RestClient
	verifyServiceSid string
}

func NewTwilioAdapter(client *twilio.RestClient, verifyServiceSid string) *TwilioAdapter {
	return &TwilioAdapter{
		client:           client,
		verifyServiceSid: verifyServiceSid,
	}
}

func (v *TwilioAdapter) SendOTP(to string) error {
	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	resp, err := v.client.VerifyV2.CreateVerification(v.verifyServiceSid, params)

	if err != nil {
		return err
	} else {
		fmt.Printf("Sent verification '%s'\n", *resp.Sid)
	}
	return nil
}

func (v *TwilioAdapter) VerififyOTP(to string, code string) error {
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)

	resp, err := v.client.VerifyV2.CreateVerificationCheck(v.verifyServiceSid, params)

	if err != nil {
		fmt.Println(err.Error())
	} else if *resp.Status == "approved" {
		return nil
	} else {
		return errors.New("incorrect code")
	}
	return nil
}
