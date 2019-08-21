package cultfit

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"testing"
)

func TestProvider_bookForDate(t *testing.T) {
	type fields struct {
		BaseURL          string
		ClassInCenterURL string
		logger           *logrus.Logger
	}
	type args struct {
		date        string
		classSlots  []cultClass
		preferences []SlotPreference
		cookie      string
		apiKey      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Provider{
				BaseURL:          tt.fields.BaseURL,
				ClassInCenterURL: tt.fields.ClassInCenterURL,
				logger:           tt.fields.logger,
			}
			if err := p.bookForDate(tt.args.date, tt.args.classSlots, tt.args.preferences, tt.args.cookie, tt.args.apiKey); (err != nil) != tt.wantErr {
				t.Errorf("bookForDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProvider_bookForDate2(t *testing.T) {
	logger := logrus.New()
	logger.Out = ioutil.Discard

	provider := Provider{
		BaseURL:          "",
		ClassInCenterURL: "",
		logger:           logger,
	}
}
