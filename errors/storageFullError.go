package errors

const (
	error_text_storage = "storage is full"
)

type StorageFullError struct{}

func (s *StorageFullError) Error() string {
	return error_text_storage
}

func NewStorageFullError() StorageFullError {
	return StorageFullError{}
}

func IsStorageFullError(err error) bool {
	if err == nil {
		return false
	}

	return err.Error() == error_text_storage
}
