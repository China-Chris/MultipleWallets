package msg

import (
	"Multiplewallets/errorss/errors_const"
)

// ErrorCodeTextMapEn 定义英文的业务报错信息
var ErrorCodeTextMapEn = map[int]string{
	errors_const.ErrInternalServer:     "Internal server error",
	errors_const.ErrorCodeInvalidInput: "Code Invalid Input",
}
