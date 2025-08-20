package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
)

type twilioConfig struct {
	AccountSid        string
	AuthToken         string
	TwilioPhoneNumber string
	TargetNumber      string
	Client            *twilio.RestClient
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading environment variables", err)
		return
	}

	cfg := twilioConfig{
		AccountSid:        os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken:         os.Getenv("TWILIO_AUTH_TOKEN"),
		TwilioPhoneNumber: os.Getenv("TWILIO_PHONE_NUMBER"),
		TargetNumber:      os.Getenv("YOUR_PHONE_NUMBER"),
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.AccountSid,
		Password: cfg.AuthToken,
	})

	cfg.Client = client

	err = cfg.sendMessage("hello from go! iss a me")
	if err != nil {
		fmt.Println("Error sending message:", err.Error())
	} else {
		fmt.Println("Message sent successfully!")
	}
}
