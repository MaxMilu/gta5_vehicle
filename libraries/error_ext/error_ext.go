package error_ext

import "github.com/pkg/errors"

var checkErrorTemplate = &checkError{}

type checkError struct {
	error
}

func (e checkError) Error() string {
	return e.error.Error()
}

func New(errorMessage string) error {
	return &checkError{error: errors.New(errorMessage)}
}

// CheckError 判断错误类型是否是校验错误
func CheckError(err error) bool {
	if err == nil {
		return false
	}
	return errors.As(err, &checkErrorTemplate)
}
