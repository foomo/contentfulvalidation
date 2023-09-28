package constants

const (
	SeverityFatal Severity = "fatal"
	SeverityError Severity = "error"
	SeverityWarn  Severity = "warn"
	SeverityInfo  Severity = "info"
)

const (
	HealthError Health = "error"
	HealthWarn  Health = "warn"
	HealthOk    Health = "ok"
)

const (
	MissingQuerryFieldValues QueryError = "Missing field values"
	QuerryValueExpired       QueryError = "Querry field value is expired"
	MissingQuerryCondition   QueryError = "Missing querry condition"
	MissingQuerryField       QueryError = "Querry Field is empty"
)

const DateFormat = "02 Jan 2006"
