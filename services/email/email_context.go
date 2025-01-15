package email

type EmailContext struct {
	Sender           string
	Receivers        []string
	Subject          string
	Languages        []string
	TemplateLocation string
	Context          map[string]interface{}
}

func NewEmailContect(sender string, receivers []string, subject string, languages []string, templateLocation string) *EmailContext {
	return &EmailContext{
		Sender:           sender,
		Receivers:        receivers,
		Subject:          subject,
		Languages:        languages,
		TemplateLocation: templateLocation,
		Context:          make(map[string]interface{}),
	}
}

func (ec *EmailContext) AddContextVariable(key string, value interface{}) map[string]interface{} {
	if _, exists := ec.Context[key]; !exists {
		ec.Context[key] = value
	}
	return ec.Context
}
