package cultfit

import (
	"sync"
	"time"
)

const daysToBookForInAdvance = 7 // books for 7 days in advance

func (p Provider) AutoBook(preferences []SlotPreference, cookie, apiKey string) {
	uniqCenterIDs := getUniqueCenterIDs(preferences)
	classSlots := p.getClassSlots(uniqCenterIDs, cookie, apiKey)

	dates := make([]string, daysToBookForInAdvance)
	for i := 0; i < daysToBookForInAdvance; i++ {
		dates[i] = time.Now().AddDate(0, 0, i+1).Format("2006-01-02")
	}

	for _, date := range dates {
		p.bookForDate(date, classSlots, preferences, cookie, apiKey)
	}
}

func (p Provider) bookForDate(date string, classSlots []cultClass, preferences []SlotPreference, cookie, apiKey string) {
	availableClassesForDay := make([]cultClass, 0)
	for _, class := range classSlots {
		if class.Date == date && (class.State == "AVAILABLE" || class.State == "BOOKED") {
			availableClassesForDay = append(availableClassesForDay, class)
			if class.State == "BOOKED" {
				// class is already booked our work here is done!
				return
			}
		}
	}

	// if the execution comes here, it means we haven't booked a class for this day
	for _, pref := range preferences {
		for _, class := range availableClassesForDay {
			if pref.CenterID == class.CenterID && pref.Time == class.StartTime && pref.WorkoutName == class.WorkoutName {
				_, err := p.BookClass(class)
				if err != nil {
					return
				}
				// error booking the class, let's move on to other preference
			}
		}
	}
}

func (p Provider) getClassSlots(centerIDs []int, cookie, apiKey string) []cultClass {
	outputCh := make(chan []cultClass)
	wg := sync.WaitGroup{}
	wg.Add(len(centerIDs))
	for _, centerID := range centerIDs {
		go func(outputCh chan<- []cultClass, centerID int) {
			defer wg.Done()
			response := <-p.FetchClassesInCenter(centerID, cookie, apiKey)
			if response.Err == nil && len(response.Data) > 0 {
				outputCh <- response.Data
			}
		}(outputCh, centerID)
	}

	go func() {
		wg.Wait()
		close(outputCh)
	}()

	allAvailableClasses := make([]cultClass, 0)
	for classes := range outputCh {
		allAvailableClasses = append(allAvailableClasses, classes...)
	}
	return allAvailableClasses
}

func getUniqueCenterIDs(preferences []SlotPreference) []int {
	centerIDsMap := make(map[int]struct{})
	for _, pref := range preferences {
		centerIDsMap[pref.CenterID] = struct{}{}
	}

	uniqCenterIDs := make([]int, 0, len(centerIDsMap))
	for centerID := range centerIDsMap {
		uniqCenterIDs = append(uniqCenterIDs, centerID)
	}
	return uniqCenterIDs
}
