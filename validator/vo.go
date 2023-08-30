package validator

type SysType string
type ModelType string
type ModelID string
type MessageCode string
type Severity string
type Health string

type ValidationResults map[ModelType]map[ModelID]*ValidationResult

type ModelTypeInfo struct {
	ModelType ModelType `json:"modelType"`
	Title     string    `json:"title"`
}
