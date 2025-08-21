package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
)

type Config struct {
	AccountSid         string
	AuthToken          string
	TwilioPhoneNumber  string
	TargetNumber       string
	StravaClientID     string
	StravaClientSecret string
	StravaRefreshToken string
	StravaAccessToken  string
	StravaAtheleteID   string
	StravaRunYearGoal  string
	Client             *twilio.RestClient
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading environment variables, assuming default has been set")
	}

	cfg := Config{
		AccountSid:         os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken:          os.Getenv("TWILIO_AUTH_TOKEN"),
		TwilioPhoneNumber:  os.Getenv("TWILIO_PHONE_NUMBER"),
		TargetNumber:       os.Getenv("YOUR_PHONE_NUMBER"),
		StravaClientID:     os.Getenv("CLIENT_ID"),
		StravaClientSecret: os.Getenv("CLIENT_SECRET"),
		StravaRefreshToken: os.Getenv("REFRESH_TOKEN"),
		StravaAtheleteID:   os.Getenv("ATHELETE_ID"),
		StravaRunYearGoal:  os.Getenv("RUN_YEAR_GOAL"),
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.AccountSid,
		Password: cfg.AuthToken,
	})

	cfg.Client = client

	err = cfg.getAccessToken()
	if err != nil {
		fmt.Println("Error: could not get strava access token", err)
		return
	}

	distance, err := cfg.getYTDDistance()
	if err != nil {
		fmt.Println("Error: could not get strava running distance", err)
		return
	}

	message, err := cfg.messageDistance(distance)
	err = cfg.sendMessage(message)
	if err != nil {
		fmt.Println("Error sending message:", err.Error())
	} else {
		fmt.Println("Message sent successfully!")
		fmt.Println("Message:")
		fmt.Println(message)
	}
}
