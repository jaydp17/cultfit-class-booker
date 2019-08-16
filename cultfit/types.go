package cultfit

type cultDay struct {
	Id    string `json:"id"`    // eg. "2019-08-17"
	Day   string `json:"day"`   // eg. "17"
	Month string `json:"month"` // eg. "Aug"
}

type cultWorkout struct {
	Id          int    `json:"id"`          // eg. 8
	Name        string `json:"name"`        // eg. "Boxing",
	DisplayText string `json:"displayText"` // eg. "Boxing"
}

type cultClass struct {
	Id             string `json:"id"`             // eg. "1034146"
	ProductType    string `json:"productType"`    // eg. "FITNESS"
	Date           string `json:"date"`           // eg. "2019-08-17"
	StartTime      string `json:"startTime"`      // eg. "06:00:00"
	EndTime        string `json:"endTime"`        // eg. "06:50:00"
	WorkoutID      int    `json:"workoutId"`      // eg.  58
	CenterID       int    `json:"centerID"`       // eg. 172
	AvailableSeats int    `json:"availableSeats"` // eg. 2
	WorkoutName    string `json:"workoutName"`    // eg. "LEARN TO SWIM"
	State          string `json:"state"`          // eg. "SEAT_NOT_AVAILABLE" or "AVAILABLE"
}
