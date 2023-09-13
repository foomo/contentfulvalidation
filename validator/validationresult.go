package validator

import "github.com/foomo/contentfulvalidation/constants"

type ValidationResult struct {
	ID              ModelID                    `json:"id"`
	Title           string                     `json:"title,omitempty"`
	InternalTitle   string                     `json:"internalTitle,omitempty"`
	LastUpdatedDate string                     `json:"lastUpdatedDate,omitempty"`
	ModelType       ModelType                  `json:"modelType"`
	Health          constants.Health           `json:"health"`
	Messages        []*ValidationResultMessage `json:"messages"`
}

type ValidationResultMessage struct {
	Code     MessageCode        `json:"code"`
	Message  string             `json:"message"`
	Severity constants.Severity `json:"severity"`
}

func (result *ValidationResult) Log(severity constants.Severity, message string, code MessageCode) {
	msg := &ValidationResultMessage{
		Code:     code,
		Message:  message,
		Severity: severity,
	}
	result.Messages = append(result.Messages, msg)
}

func (result *ValidationResult) UpdateHealth() {
	if len(result.Messages) > 0 {
		for _, msg := range result.Messages {
			if msg.Severity == constants.SeverityError || msg.Severity == constants.SeverityFatal {
				result.Health = constants.HealthError
				return
			}
			if msg.Severity == constants.SeverityWarn {
				result.Health = constants.HealthWarn
			}
		}
	}
}
