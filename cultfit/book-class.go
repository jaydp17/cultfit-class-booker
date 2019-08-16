package cultfit

import (
	"fmt"
	"github.com/imroc/req"
	"net/http"
)

type BookClassResult struct {
	Booked bool
	Err    error
}

func (p Provider) BookClass(class cultClass, cookie, apiKey string) <-chan BookClassResult {
	resultCh := make(chan BookClassResult)

	go func() {
		defer close(resultCh)
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
		res, err := req.Get(p.getClassBookingURL(class.ID), headers)
		if err != nil {
			resultCh <- BookClassResult{Err: fmt.Errorf("failed to book class: %v", err)}
			return
		}

		if res.Response().StatusCode != http.StatusOK {
			resultCh <- BookClassResult{Err: fmt.Errorf("failed to book class: %v", res.Response().Status)}
			return
		}

		resultCh <- BookClassResult{Booked: true}
	}()

	return resultCh
}
