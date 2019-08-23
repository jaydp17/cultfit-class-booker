package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jaydp17/cultfit-class-booker/cultfit"
	"github.com/jaydp17/cultfit-class-booker/logger"
)

const swimmingHSRCenterID = 172
const swimmingKoramangalaCenterID = 164

func Handler() {
	log := logger.New()
	log.Info("starting handler")
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
			CenterID:    swimmingKoramangalaCenterID,
			Time:        "18:00:00", // 6pm
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
	log.Info("done!")
}

func main() {
	lambda.Start(Handler)
}
