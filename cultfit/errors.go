package cultfit

import "fmt"

type NoAvailableClassError struct {
	Date string
}

func (e NoAvailableClassError) Error() string {
	return fmt.Sprintf("No available class found for date: %s", e.Date)
}

type NoMatchingClassError struct {
	otherAvailableSlots map[int][]string // <CenterID, timings>
}

func (e NoMatchingClassError) Error() string {
	return fmt.Sprintf("No matching class found, however there are other slots for the same workout in the same centers: %+v", e.otherAvailableSlots)
}
