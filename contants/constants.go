package contants

import (
	"github.com/foomo/contentfulvalidation/validator"
)

const (
	SeverityFatal validator.Severity = "fatal"
	SeverityError validator.Severity = "error"
	SeverityWarn  validator.Severity = "warn"
	SeverityInfo  validator.Severity = "info"
)

const (
	HealthError validator.Health = "error"
	HealthWarn  validator.Health = "warn"
	HealthOk    validator.Health = "ok"
)

const DateFormat = "02 Jan 2006"
