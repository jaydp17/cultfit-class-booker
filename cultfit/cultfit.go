package cultfit

import (
	"fmt"
	"github.com/jaydp17/cultfit-class-booker/logger"
	"github.com/sirupsen/logrus"
)

type Provider struct {
	BaseURL          string
	ClassInCenterURL string
	logger           *logrus.Logger
}

func (p Provider) getClassBookingURL(classID string) string {
	return fmt.Sprintf("%s/class/%s/book", p.BaseURL, classID)
}

func New() Provider {
	cultProvider := Provider{
		BaseURL: "https://www.cure.fit/api/cult",
		logger:  logger.New(),
	}
	cultProvider.ClassInCenterURL = cultProvider.BaseURL + "/classes"
	return cultProvider
}
