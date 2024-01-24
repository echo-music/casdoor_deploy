package exception

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const defaultCode = -1

type ApiError struct {
	code int
	err  error
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("Error with code %d: %s", e.code, e.err.Error())
}

func New(message string) error {
	return &ApiError{
		code: defaultCode,
		err:  errors.New(message),
	}
}

func WithMessage(err error, message string) error {
	return &ApiError{
		code: defaultCode,
		err:  errors.WithMessage(err, message),
	}
}
func WithCodeAndMessage(code int, err error, message string) error {
	return &ApiError{
		code: code,
		err:  errors.WithMessage(err, message),
	}
}

func WithStack(err error) error {
	return &ApiError{
		code: defaultCode,
		err:  errors.WithStack(err),
	}
}

func RealErr(err error) error {
	fmt.Println("start")
	if targetErr := new(ApiError); errors.As(err, &targetErr) {
		fmt.Println("ApiError")
		return targetErr.err
	}
	if targetErr := new(gin.Error); errors.As(err, &targetErr) {
		fmt.Println("gin.Error")
		return targetErr.Err
	}
	return err
}

// 获取错误码
func Code(err error) int {
	if e := new(ApiError); errors.As(err, &e) {
		return e.code
	}
	return defaultCode
}
