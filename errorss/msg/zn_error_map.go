package msg

import "Multiplewallets/errorss/errors_const"

// ErrorCodeTextMapZh 定义中文的业务报错信息
var ErrorCodeTextMapZh = map[int]string{
	errors_const.ErrInternalServer:            "内部服务发生错误",
	errors_const.ErrorCodeInvalidInput:        "无效输入参数",
	errors_const.ErrorInvalidWalletAddress:    "无效钱包地址",
	errors_const.ErrorInvalidAddress:          "无效用户地址",
	errors_const.ErrorDuplicateWalletCreation: "重复创建钱包",
}
