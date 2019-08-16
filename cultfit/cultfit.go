package cultfit

type Provider struct {
	BaseURL          string
	ClassInCenterURL string
}

func New() Provider {
	cultProvider := Provider{
		BaseURL: "https://www.cure.fit/api/cult",
	}
	cultProvider.ClassInCenterURL = cultProvider.BaseURL + "/classes"
	return cultProvider
}
