package cultfit

import (
	"fmt"
	"github.com/imroc/req"
)

type FetchClassInCenterResult struct {
	Data classInCenterResponse
	Err  error
}

func (p Provider) FetchClassesInCenter(centerID int, cookie, apiKey string) <-chan FetchClassInCenterResult {
	resultCh := make(chan FetchClassInCenterResult)

	go func() {
		defer close(resultCh)
		params := req.QueryParam{"center": centerID}
		headers := req.Header{
			"sec-fetch-mode": "cors",
			"osname":         "browser",
			"cookie":         cookie,
			"pragma":         "no-cache",
			"user-agent":     "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Mobile Safari/537.36",
			"content-type":   "application/json",
			"accept":         "application/json",
			"cache-control":  "no-cache,no-cache",
			"authority":      "www.cure.fit",
			"apikey":         apiKey,
			"sec-fetch-site": "same-origin",
			"referer":        "https://www.cure.fit/cult/classbooking?pageFrom=cultCLP&pageType=classbooking",
			"appversion":     "7",
		}

		res, err := req.Get(p.ClassInCenterURL, params, headers)
		if err != nil {
			resultCh <- FetchClassInCenterResult{Err: fmt.Errorf("failed fetching classes in a center: %v", err)}
			return
		}

		var jsonResp classInCenterResponse
		if err := res.ToJSON(&jsonResp); err != nil {
			resultCh <- FetchClassInCenterResult{Err: fmt.Errorf("failed fetching classes in a center: %v", err)}
			return
		}

		resultCh <- FetchClassInCenterResult{Data: jsonResp}
	}()
	return resultCh
}

type classInCenterResponse struct {
	Title           string                 `json:"title"`
	Days            []cultDay              `json:"days"`
	WorkoutFilters  []cultWorkout          `json:"workoutFilters"`
	ClassByDateList []classesGroupedByDate `json:"classByDateList"`
}

type classesGroupedByDate struct {
	Id              string                 `json:"id"`         // eg. "2019-08-17"
	WidgetType      string                 `json:"widgetType"` // eg. "BROWSE_CLASS_LIST"
	ClassByTimeList []classesGroupedByTime `json:"classByTimeList"`
}

type classesGroupedByTime struct {
	Id            string      `json:"id"`           // eg. "06:00:00"
	DisableGroup  bool        `json:"disableGroup"` // eg. false
	Classes       []cultClass `json:"classes"`
	NearByCenters []nearbyCenter
}

type nearbyCenter struct {
	CenterID   int         `json:"centerId"`   // eg. 3
	CenterName string      `json:"centerName"` // eg. "Cult HSR"
	Classes    []cultClass `json:"classes"`
}
