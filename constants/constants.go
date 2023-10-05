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
	MissingQueryFieldValues QueryError = "Missing field values"
	QueryValueExpired       QueryError = "Query field value is expired"
	MissingQueryCondition   QueryError = "Missing query condition"
	MissingQueryField       QueryError = "Query Field is empty"
)

const DateFormat = "02 Jan 2006"
