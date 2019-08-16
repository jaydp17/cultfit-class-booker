package main

import (
	"github.com/jaydp17/cultfit-class-booker/cultfit"
)

const swimmingHSRCenterID = 172
const swimmingKoramangalaCenterID = 164

func main() {
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
