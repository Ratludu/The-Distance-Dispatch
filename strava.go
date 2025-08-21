package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const OAuthURL = "https://www.strava.com/api/v3/oauth/token"

type StravaToken struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpiresAt    int    `json:"expires_at"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type StravaData struct {
	RecentRunTotals struct {
		Count            int     `json:"count"`
		Distance         float64 `json:"distance"`
		MovingTime       int     `json:"moving_time"`
		ElapsedTime      int     `json:"elapsed_time"`
		ElevationGain    float64 `json:"elevation_gain"`
		AchievementCount int     `json:"achievement_count"`
	} `json:"recent_run_totals"`
	AllRunTotals struct {
		Count         int     `json:"count"`
		Distance      float64 `json:"distance"`
		MovingTime    int     `json:"moving_time"`
		ElapsedTime   int     `json:"elapsed_time"`
		ElevationGain float64 `json:"elevation_gain"`
	} `json:"all_run_totals"`
	YtdRunTotals struct {
		Count         int     `json:"count"`
		Distance      int     `json:"distance"`
		MovingTime    float64 `json:"moving_time"`
		ElapsedTime   float64 `json:"elapsed_time"`
		ElevationGain float64 `json:"elevation_gain"`
	} `json:"ytd_run_totals"`
}

func (c *Config) getAccessToken() error {

	data := url.Values{}
	data.Set("client_id", c.StravaClientID)
	data.Set("client_secret", c.StravaClientSecret)
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", c.StravaRefreshToken)

	req, err := http.NewRequest("POST", OAuthURL, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("Error: Could not create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error doing request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error: Could not read body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Error: Could not get access token, api failed with error code %d", resp.StatusCode)
	}

	var stravaToken StravaToken
	if err = json.Unmarshal(body, &stravaToken); err != nil {
		return fmt.Errorf("Error: Failed to unmarshal strava token data: %v", err)
	}

	c.StravaAccessToken = stravaToken.AccessToken

	return nil
}

func (c *Config) getYTDDistance() (int, error) {

	statsURL := fmt.Sprintf("https://www.strava.com/api/v3/athletes/%s/stats", c.StravaAtheleteID)

	req, err := http.NewRequest("GET", statsURL, nil)
	if err != nil {
		return 0, fmt.Errorf("Error: Could not create request: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.StravaAccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("Error doing request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("Error: Could not read body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Error: Could not get access token, api failed with error code %d", resp.StatusCode)
	}

	var stravaData StravaData
	if err = json.Unmarshal(body, &stravaData); err != nil {
		return 0, fmt.Errorf("Error: Failed to unmarshal strava data: %v", err)
	}

	return stravaData.YtdRunTotals.Distance, nil
}

func (c *Config) messageDistance(distance int) (string, error) {

	goal, err := strconv.Atoi(c.StravaRunYearGoal)
	if err != nil {
		return "", err
	}

	distanceKm := float64(distance) / 1000.0
	percentage := (float64(distanceKm) / float64(goal)) * 100

	remainingDays := getRemainingDays()
	averageKMremaining := (float64(goal) - float64(distanceKm)) / float64(remainingDays)

	return fmt.Sprintf("The distance dispatch - beep boop\n- Progress Update: %.2fkm of %d (%.2f%%)\n- Average Km/day remaining: %.2fkm (%d days)", distanceKm, goal, percentage, averageKMremaining, remainingDays), nil
}

func getRemainingDays() int {
	now := time.Now()
	currentDay := now.YearDay()
	return 365 - currentDay
}
