package errorss

import (
	"Multiplewallets/errorss/msg"
	"errors"

	"fmt"
)

// CustomError 包装自定义错误信息
type CustomError struct {
	Code    int
	Message string
	Err     error
}

// Error 返回错误字符串
func (ce *CustomError) Error() string {
	return fmt.Sprintf("code:%d, message:%s, err:%s", ce.Code, ce.Message, ce.Err)
}

// HandleError 统一错误处理程序，支持多语言（默认中文）
func HandleError(code int, language string, err error) *CustomError {
	var errorCodeTextMap map[int]string
	switch language {
	case "en":
		errorCodeTextMap = msg.ErrorCodeTextMapEn
	//case "ja":
	//	errorCodeTextMap = msg.err
	//case "ko":
	//	errorCodeTextMap = msg.ErrorCodeTextMapKo
	default:
		errorCodeTextMap = msg.ErrorCodeTextMapZh
	}

	return &CustomError{
		Code:    code,
		Message: errorCodeTextMap[code],
		Err:     err,
	}
}

// HandleErrorZh 中文错误处理
func HandleErrorZh(code int, err error) *CustomError {
	return HandleError(code, "zh", err)
}

// HandleErrorEn 英文错误处理
func HandleErrorEn(code int, err error) *CustomError {
	return HandleError(code, "en", err)
}

// HandleErrorJa 日语错误处理
func HandleErrorJa(code int, err error) *CustomError {
	return HandleError(code, "ja", err)
}

// HandleErrorKo 韩语错误处理
func HandleErrorKo(code int, err error) *CustomError {
	return HandleError(code, "ko", err)
}

// NewCustomErrorWithCode 创建自定义错误实例，使用指定错误代码和语言
func NewCustomErrorWithCode(code int, language string) error {
	errorMessage := msg.ErrorCodeTextMapZh[code]
	return errors.New(errorMessage)
}
