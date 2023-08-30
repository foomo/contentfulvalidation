package errors

const (
	ValidationErrorUnknownType       ValidationError = "UNKNOWN_MODEL_TYPE"
	ValidationErrorContentfulLoading ValidationError = "CONTENTFUL_LOADING_ERROR"
)

var (
	ErrValidationErrorUnknownType       = NewError(ValidationErrorUnknownType)
	ErrValidationErrorContentfulLoading = NewError(ValidationErrorContentfulLoading)
)

type (
	ValidationError string
)

// ValidationError constructor
func NewError(e ValidationError) *ValidationError {
	return &e
}

// Error api method
func (e *ValidationError) Error() string {
	if e != nil {
		return string(*e)
	}
	return ""
}
