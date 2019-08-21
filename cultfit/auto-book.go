package cultfit

import (
	"github.com/sirupsen/logrus"
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
		if err := p.bookForDate(date, classSlots, preferences, cookie, apiKey); err != nil {
			switch errorT := err.(type) {

			case NoAvailableClassError:
				p.logger.WithFields(logrus.Fields{
					"date": date,
				}).Info("no available classes")

			case NoMatchingClassError:
				p.logger.WithFields(logrus.Fields{
					"preferredCenterSlots": errorT.otherAvailableSlots,
				}).Error("failed to find a matching slot")
			}

		}
	}
}

func (p Provider) bookForDate(date string, classSlots []cultClass, preferences []SlotPreference, cookie, apiKey string) error {
	availableClassesForDay := make([]cultClass, 0)
	for _, class := range classSlots {
		if class.Date == date && (class.State == "AVAILABLE" || class.State == "BOOKED") {
			availableClassesForDay = append(availableClassesForDay, class)
			// RETURN early if class already booked
			if class.State == "BOOKED" {
				p.logger.WithFields(logrus.Fields{
					"date":  date,
					"class": class,
				}).Info("class already booked")
				return nil
			}
		}
	}

	if len(availableClassesForDay) == 0 {
		return NoAvailableClassError{date}
	}

	// TODO: log state of each preference

	foundPreferredSlot := false
	// if the execution comes here, it means we haven't booked a class for this day
	for _, pref := range preferences {
		for _, class := range availableClassesForDay {
			if pref.CenterID == class.CenterID && pref.Time == class.StartTime && pref.WorkoutName == class.WorkoutName {
				foundPreferredSlot = true
				bookingResult := <-p.BookClass(class, cookie, apiKey)
				if bookingResult.Booked {
					p.logger.WithFields(logrus.Fields{
						"date":  date,
						"class": class,
					}).Info("successfully booked a class")
					// class booked, let's not go over other preferences
					return nil
				}

				if bookingResult.Err != nil {
					p.logger.WithFields(logrus.Fields{
						"date":         date,
						"class":        class,
						"bookingError": bookingResult.Err.Error(),
					}).Error("error booking the class")
				}
				// error booking the class, let's move on to other preference
			}
		}
	}

	if !foundPreferredSlot {
		// stores <centerID, slot time>
		preferredCenterSlots := make(map[int][]string)
		for _, pref := range preferences {
			for _, class := range availableClassesForDay {
				if pref.CenterID == class.CenterID && pref.WorkoutName == class.WorkoutName {
					if slots, ok := preferredCenterSlots[pref.CenterID]; ok {
						slots = append(slots, class.StartTime)
						continue
					}
					preferredCenterSlots[pref.CenterID] = []string{class.StartTime}
				}
			}
		}
		return NoMatchingClassError{preferredCenterSlots}
	}

	return nil
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
			if response.Err != nil {
				p.logger.WithFields(logrus.Fields{
					"centerID": centerID,
					"response": response,
				}).Error("error fetching classes in a center")
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
