package handle

import (
	"Multiplewallets/errorss"
	"Multiplewallets/errorss/errors_const"
	"Multiplewallets/request"
	"Multiplewallets/response"
	"Multiplewallets/service"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

// StringError 是一个自定义错误类型，用于将字符串转换为错误类型。
type StringError string

// Error 返回字符串错误的字符串表示形式。
func (e StringError) Error() string {
	return string(e)
}

// 自定义验证函数，检查是否为有效的0x地址
func isValidAddress(fl validator.FieldLevel) bool {
	address := fl.Field().String()
	return common.IsHexAddress(address)
}

var validate = validator.New()

func init() {
	// 注册自定义验证函数并检查是否注册成功
	err := validate.RegisterValidation("isValidAddress", isValidAddress)
	if err != nil {
		return
	}
}

// handleErrorResponse 函数用于处理错误响应并中断请求。
func handleErrorResponse(c *gin.Context, statusCode int, errorCode int, err error) {
	c.AbortWithStatusJSON(statusCode, errorss.HandleError(errorCode, "zh", StringError(err.Error())))
}

// validateVarAddress 函数用于验证给定的地址是否是有效的 0x 地址。
// 如果地址无效，则返回相应的错误；否则，返回 nil 表示地址有效。
func validateVarAddress(address string) error {
	if err := validate.Var(address, "isValidAddress"); err != nil {
		return StringError(err.Error())
	}
	return nil
}

// CreateMultipleSignatureWallet 创建多重签名钱包
// @Description 创建具有指定地址的多重签名钱包。
// @Tags 钱包
// @Accept json
// @Produce json
// @Param request body  request.CreateMultipleWallet true "用于创建多重签名钱包的请求体"
// @Success 201 {object} response.CreateMultipleWalletResponse
// @Router /create-multiple-wallet [post]
func CreateMultipleSignatureWallet(c *gin.Context) {
	req, err := request.BindCreateMultipleWallet(c)
	if err != nil {
		handleErrorResponse(c, http.StatusBadRequest, errors_const.ErrInternalServer, err)
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validateVarAddress(req.Address); err != nil {
		handleErrorResponse(c, http.StatusBadRequest, errors_const.ErrorInvalidWalletAddress, err)
		return
	}
	// 调用服务层的函数处理业务逻辑
	createdAddress, err := service.CreateMultipleWallet(req)
	if err != nil {
		handleErrorResponse(c, http.StatusInternalServerError, errors_const.ErrInternalServer, err)
		return
	}
	// 根据服务层返回的创建地址信息组装响应结构体
	resp := response.CreateMultipleWalletResponse{
		Address: createdAddress,
	}
	// 201代表资源创建成功
	c.JSON(http.StatusOK, resp)
}

// AddMembers 添加成员
func AddMembers(c *gin.Context) {
	req, err := request.BindAddMembers(c)
	if err != nil {
		handleErrorResponse(c, http.StatusBadRequest, errors_const.ErrInternalServer, err)
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validateVarAddress(req.WalletAddress); err != nil {
		handleErrorResponse(c, http.StatusBadRequest, errors_const.ErrorInvalidWalletAddress, err)
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validateVarAddress(req.Address); err != nil {
		handleErrorResponse(c, http.StatusBadRequest, errors_const.ErrorInvalidAddress, err)
		return
	}
	// 调用服务层的函数处理业务逻辑
	bool, err := service.AddMembers(req)
	if err != nil {
		handleErrorResponse(c, http.StatusInternalServerError, errors_const.ErrInternalServer, err)
		return
	}
	resp := response.AddMembersResponse{
		Bool: bool,
	}
	c.JSON(http.StatusOK, resp)
}

// AddWeight 添加权重成员
func AddWeight(c *gin.Context) {
	req, err := request.BinAddWeight(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrInternalServer, "zh", err))
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.WalletAddress, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidWalletAddress, "zh", StringError(err.Error())))
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.Address, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}
	bool, err := service.AddWeight(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}
	resp := response.AddWeightResponse{
		Bool: bool,
	}
	c.JSON(http.StatusOK, resp)
}

// TxTransCation 事务交易
func TxTransCation(c *gin.Context) {
	req, err := request.BinTxTransCation(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrInternalServer, "zh", err))
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.WalletAddress, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidWalletAddress, "zh", StringError(err.Error())))
		return
	}
	bool, err := service.CreateTxTransCation(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}
	resp := response.CreateTxTransCationResponse{
		Bool: bool,
	}
	c.JSON(http.StatusOK, resp)
}

func NewTransCationNumber(c *gin.Context) {
	req, err := request.BinNewTransCationNumber(c)
	if err != nil {
		fmt.Println(3213)
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrInternalServer, "zh", err))
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.WalletAddress, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidWalletAddress, "zh", StringError(err.Error())))
		return
	}
	nonce, err := service.GetNewTransCationNumber(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}
	resp := response.NewTransCationNumberResponse{
		Nonce: nonce,
	}
	c.JSON(http.StatusOK, resp)
}

func SignTxTransCation(c *gin.Context) {
	req, err := request.BindSignTxTransCation(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrInternalServer, "zh", err))
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.WalletAddress, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidWalletAddress, "zh", StringError(err.Error())))
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.Address, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}
	bool, err := service.SignTxTransCation(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}
	resp := response.SignTxTransCationResponse{
		Bool: bool,
	}
	c.JSON(http.StatusOK, resp)
}

func VerifyTransaction(c *gin.Context) {
	req, err := request.BindVerifyTransaction(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrInternalServer, "zh", err))
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.WalletAddress, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidWalletAddress, "zh", StringError(err.Error())))
		return
	}
	list, err := service.VerifyTransaction(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}
	c.JSON(http.StatusOK, list)
}

func CancelTransaction(c *gin.Context) {
	req, err := request.BindCancelTransaction(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrInternalServer, "zh", err))
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.WalletAddress, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidWalletAddress, "zh", StringError(err.Error())))
		return
	}
	bool, err := service.CancelTransaction(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}
	resp := response.CancelTransactionResponse{
		Bool: bool,
	}
	c.JSON(http.StatusOK, resp)
}

func TransactionList(c *gin.Context) {
	// 从查询参数中获取钱包地址
	walletAddress := c.Query("wallet_address")
	nonceStr := c.Query("nonce")
	fmt.Println(nonceStr)
	nonce, err := strconv.Atoi(nonceStr)
	if err != nil {
		// 处理非整数值的情况，例如返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrInternalServer, "zh", nil))
		return
	}
	// 构建请求对象
	req := request.TransactionList{
		WalletAddress: walletAddress,
		Nonce:         nonce,
	}
	// 检查获取的地址是否有效
	if err := validate.Var(walletAddress, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidWalletAddress, "zh", StringError(err.Error())))
		return
	}

	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(walletAddress, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidWalletAddress, "zh", StringError(err.Error())))
		return
	}
	signedAddresses, unsignedAddresses, err := service.TransactionLists(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}
	resp := response.TransactionListResponse{
		SignedAddresses:   signedAddresses,
		UnsignedAddresses: unsignedAddresses,
	}

	c.JSON(http.StatusOK, resp)
}

func TxCompleted(c *gin.Context) {
	req, err := request.BindTxCompleted(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrInternalServer, "zh", err))
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.WalletAddress, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidWalletAddress, "zh", StringError(err.Error())))
		return
	}
	bool, err := service.TxCompleted(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}
	resp := response.TxCompletedResponse{Bool: bool}

	c.JSON(http.StatusOK, resp)
}

func UpdateThreshold(c *gin.Context) {
	req, err := request.BinUpdateThreshold(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrInternalServer, "zh", err))
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.WalletAddress, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidWalletAddress, "zh", StringError(err.Error())))
		return
	}

	bool, err := service.UpdateThreshold(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}
	resp := response.UpdateThresholdResponse{Bool: bool}

	c.JSON(http.StatusOK, resp)
}

func UpdateWeight(c *gin.Context) {
	req, err := request.BinUpdateWeight(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrInternalServer, "zh", err))
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.WalletAddress, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidWalletAddress, "zh", StringError(err.Error())))
		return
	}
	bool, err := service.UpdateWeight(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}
	//resp := response.UpdateWeightResponse{Bool: bool}
	c.JSON(http.StatusOK, bool)
}

func GetUserInfo(c *gin.Context) {
	req, err := request.BinGetUserInfo(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrInternalServer, "zh", err))
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.WalletAddress, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidWalletAddress, "zh", StringError(err.Error())))
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.Address, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}
	name, img, err := service.GetUserInfo(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}
	resp := response.GetUserInfoResponse{
		Name: name,
		Img:  img,
	}
	c.JSON(http.StatusOK, resp)
}

func VerifyTransactionBeReady(c *gin.Context) {
	req, err := request.BinVerifyTransactionBeReady(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrInternalServer, "zh", err))
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.WalletAddress, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidWalletAddress, "zh", StringError(err.Error())))
		return
	}
	bool, err := service.VerifyTransactionBeReady(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}
	resp := response.VerifyTransactionBeReady{
		Bool: bool,
	}
	c.JSON(http.StatusOK, resp)
}

func TransactionHistory(c *gin.Context) {
	//func TransactionList(c *gin.Context) {
	//	// 从查询参数中获取钱包地址
	walletAddress := c.Query("wallet_address")
	//	nonceStr := c.Query("nonce")
	//	fmt.Println(nonceStr)
	//	nonce, err := strconv.Atoi(nonceStr)
	//	if err != nil {
	//		// 处理非整数值的情况，例如返回错误响应
	//		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrInternalServer, "zh", nil))
	//		return
	//	}
	// 构建请求对象
	req := request.TransactionHistory{
		WalletAddress: walletAddress,
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.WalletAddress, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidWalletAddress, "zh", StringError(err.Error())))
		return
	}
	transactionHistory, err := service.TransactionHistory(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}
	c.JSON(http.StatusOK, transactionHistory)
}

// DeleteMember godoc
// @Summary 删除成员
// @Description 删除指定钱包地址的成员
// @Produce json
// @Param wallet_address query string true "钱包地址"
// @Param address query string true "成员地址"
// @Success 200 {boolean} bool true "删除成功"
// @Router /deleteMember [delete]
func DeleteMember(c *gin.Context) {
	walletAddress := c.Query("wallet_address")
	address := c.Query("address")
	req := request.DeleteMember{
		WalletAddress: walletAddress,
		Address:       address,
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.WalletAddress, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidWalletAddress, "zh", StringError(err.Error())))
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.Address, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidAddress, "zh", StringError(err.Error())))
		return
	}
	bool, err := service.DeleteMember(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}

	c.JSON(http.StatusOK, bool)
}

func ReplaceMemberAddress(c *gin.Context) {
	req, err := request.BinReplaceMemberAddressBeReady(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrInternalServer, "zh", err))
		return
	}
	fmt.Println("123")
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.WalletAddress, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidWalletAddress, "zh", StringError(err.Error())))
		return
	}
	// 校验req对象中的Address字段是否为有效的0x地址
	if err := validate.Var(req.Address, "isValidAddress"); err != nil {
		// 如果Address字段不是有效的0x地址，返回错误响应
		c.JSON(http.StatusBadRequest, errorss.HandleError(errors_const.ErrorInvalidWalletAddress, "zh", StringError(err.Error())))
		return
	}
	bool, err := service.ReplaceMemberAddress(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorss.HandleError(errors_const.ErrInternalServer, "zh", StringError(err.Error())))
		return
	}
	resp := response.VerifyTransactionBeReady{
		Bool: bool,
	}
	c.JSON(http.StatusOK, resp)
}
