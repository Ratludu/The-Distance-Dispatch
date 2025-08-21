package main

import (
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func (c *Config) sendMessage(msg string) error {

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(c.TargetNumber)
	params.SetFrom(c.TwilioPhoneNumber)
	params.SetBody(msg)

	_, err := c.Client.Api.CreateMessage(params)
	if err != nil {
		return err
	}
	return nil
}
