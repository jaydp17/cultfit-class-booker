package cultfit

import "fmt"

type Provider struct {
	BaseURL          string
	ClassInCenterURL string
}

func (p Provider) getClassBookingURL(classID string) string {
	return fmt.Sprintf("%s/class/%s/book", p.BaseURL, classID)
}

func New() Provider {
	cultProvider := Provider{
		BaseURL: "https://www.cure.fit/api/cult",
	}
	cultProvider.ClassInCenterURL = cultProvider.BaseURL + "/classes"
	return cultProvider
}
