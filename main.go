package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jaydp17/cultfit-class-booker/cultfit"
)

const swimmingHSRCenterID = 172
const swimmingKoramangalaCenterID = 164

func Handler() {
	cookie := getCultCookie()
	apiKey := getCultAPIKey()

	preferences := []cultfit.SlotPreference{
		{
			CenterID:    swimmingKoramangalaCenterID,
			Time:        "17:00:00", // 5pm
			WorkoutName: "LEARN TO SWIM",
		},
		{
			CenterID:    swimmingKoramangalaCenterID,
			Time:        "17:30:00", // 5:30pm
			WorkoutName: "LEARN TO SWIM",
		},
		{
			CenterID:    swimmingHSRCenterID,
			Time:        "17:00:00", // 5pm
			WorkoutName: "LEARN TO SWIM",
		},
		{
			CenterID:    swimmingHSRCenterID,
			Time:        "17:30:00", // 5:30pm
			WorkoutName: "LEARN TO SWIM",
		},
	}

	cultProvider := cultfit.New()
	cultProvider.AutoBook(preferences, cookie, apiKey)
}

func main() {
	lambda.Start(Handler)
}
