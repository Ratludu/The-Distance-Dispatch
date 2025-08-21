# The Distance Dispatch

The Distance Dispatch is a Go application that fetches your year-to-date running distance from Strava and sends you a progress update via Twilio SMS. It helps you stay on track with your annual running goals.

## Features

*   Fetches year-to-date running distance from Strava.
*   Calculates progress towards a configurable yearly running goal.
*   Sends progress updates via Twilio SMS.
*   Securely manages API keys and secrets using environment variables.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

*   Go (version 1.15 or later)
*   A Twilio account with a phone number
*   A Strava account and API application

### Installation

1.  Clone the repository:
    ```sh
    git clone https://github.com/Ratludu/The-Distance-Dispatch.git
    ```
2.  Navigate to the project directory:
    ```sh
    cd The-Distance-Dispatch
    ```
3.  Install the dependencies:
    ```sh
    go mod download
    ```

## Configuration

This project uses a `.env` file to manage environment variables. Create a file named `.env` in the root of the project and add the following variables:

```
TWILIO_ACCOUNT_SID="your_twilio_account_sid"
TWILIO_AUTH_TOKEN="your_twilio_auth_token"
TWILIO_PHONE_NUMBER="your_twilio_phone_number"
YOUR_PHONE_NUMBER="the_number_to_send_the_sms_to"
CLIENT_ID="your_strava_client_id"
CLIENT_SECRET="your_strava_client_secret"
REFRESH_TOKEN="your_strava_refresh_token"
ATHELETE_ID="your_strava_athelete_id"
RUN_YEAR_GOAL="your_yearly_running_goal_in_km"
```

## Usage (Local)

To run the application locally for testing, execute the following command:

```sh
go run .
```

This will trigger the application to fetch your Strava data and send an SMS to your configured phone number.

## Deployment (GitHub Actions)

This repository includes a GitHub Actions workflow to run the application on a schedule.

The workflow is defined in `.github/workflows/daily-dispatch.yml` and is configured to run daily at 10:00 PM AEST.

To use the GitHub Actions workflow, you need to add your environment variables as secrets to your GitHub repository.

1.  In your GitHub repository, go to `Settings` > `Secrets and variables` > `Actions`.
2.  Click `New repository secret` for each of the following secrets:
    *   `TWILIO_ACCOUNT_SID`
    *   `TWILIO_AUTH_TOKEN`
    *   `TWILIO_PHONE_NUMBER`
    *   `YOUR_PHONE_NUMBER`
    *   `CLIENT_ID`
    *   `CLIENT_SECRET`
    *   `REFRESH_TOKEN`
    *   `ATHELETE_ID`
    *   `RUN_YEAR_GOAL`

Once the secrets are added, the workflow will run automatically on the defined schedule. You can also trigger it manually from the `Actions` tab in your repository.

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1.  Fork the Project
2.  Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3.  Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4.  Push to the Branch (`git push origin feature/AmazingFeature`)
5.  Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.
