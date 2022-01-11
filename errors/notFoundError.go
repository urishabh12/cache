package errors

const (
	error_text_not_found = "not found"
)

type NotFoundError struct{}

func (s *NotFoundError) Error() string {
	return error_text_not_found
}

func NewNotFoundError() NotFoundError {
	return NotFoundError{}
}

func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	return err.Error() == error_text_not_found
}
