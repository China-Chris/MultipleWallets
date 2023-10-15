package request

import "github.com/gin-gonic/gin"

// CreateMultipleWallet 用于创建多重签名钱包的请求体
type CreateMultipleWallet struct {
	Address    string `json:"address" binding:"required"`     //地址
	Threshold  int    `json:"threshold" binding:"required"`   //多签门限值
	IsWeighted int    `json:"is_weighted" binding:"required"` //是否为权重值钱包
}

type AddMembers struct {
	WalletAddress string `json:"wallet_address"binding:"required"` //钱包地址
	Name          string `json:"name"`                             //姓名
	Img           string `json:"img"`                              //头像
	Address       string `json:"address"binding:"required"`        //添加成员地址
}

type AddWeight struct {
	WalletAddress string `json:"wallet_address"binding:"required"` //钱包地址
	Name          string `json:"name"`                             //姓名
	Img           string `json:"img"`                              //头像
	Address       string `json:"address"binding:"required"`        //添加成员地址
	Weight        int    `json:"weight"binding:"required"`         //权重值
}

type TxTransCation struct {
	WalletAddress   string `json:"wallet_address"binding:"required"`   //钱包地址
	Nonce           int    `json:"nonce"`                              // 交易序号
	TransactionType string `json:"transaction_type"binding:"required"` // 交易类型  即事务类型
	Content         string `json:"content"binding:"required"`          // 交易内容的 Json 格式
}

type NewTransCationNumber struct {
	WalletAddress string `json:"wallet_address"binding:"required"` //钱包地址
}

type SignTxTransCation struct {
	WalletAddress string `json:"wallet_address"binding:"required"` //钱包地址
	Address       string // 签名用户地址
	Nonce         int    // 交易序号
	SignatureData string `json:"signature_data""` // 签名
}

type VerifyTransaction struct {
	WalletAddress string `json:"wallet_address"binding:"required"` //钱包地址
	Nonce         int    // 交易序号
}

type CancelTransaction struct {
	WalletAddress string `json:"wallet_address"binding:"required"` //钱包地址
	Nonce         int    // 交易序号
}

type TransactionList struct {
	WalletAddress string `json:"wallet_address"binding:"required"` //钱包地址
	Nonce         int    // 交易序号
}

type TxCompleted struct {
	WalletAddress string `json:"wallet_address"binding:"required"` //钱包地址
	Nonce         int    // 交易序号
	Hash          string // 完成交易哈希
}

type UpdateThreshold struct {
	WalletAddress string `json:"wallet_address"binding:"required"` //钱包地址
	Threshold     int    `json:"threshold" binding:"required"`     //多签门限值
}

type UpdateWeight struct {
	WalletAddress string `json:"wallet_address"binding:"required"` //钱包地址
	Address       string `json:"address"binding:"required"`        //成员地址
	Weight        int    `json:"weight"binding:"required"`         //权重值
}

type GetUserInfo struct {
	WalletAddress string `json:"wallet_address"binding:"required"` //钱包地址
	Address       string `json:"address"binding:"required"`        //成员地址
	IsWeighted    int    `json:"is_weighted" binding:"required"`   //是否为权重值钱包
}

type VerifyTransactionBeReady struct {
	WalletAddress string `json:"wallet_address"binding:"required"` //钱包地址
	Nonce         int    // 交易序号
}

type TransactionHistory struct {
	WalletAddress string `json:"wallet_address"binding:"required"` //钱包地址
}

type DeleteMember struct {
	WalletAddress string `json:"wallet_address"binding:"required"` //钱包地址
	Address       string `json:"address"binding:"required"`        //删除成员地址
}

type ReplaceMemberAddress struct {
	WalletAddress string `json:"wallet_address"binding:"required"` //钱包地址
	Address       string `json:"address"binding:"required"`        //替换成员地址
	NewAddress    string `json:"new_address"binding:"required"`    //新的成员地址
}

func BindCreateMultipleWallet(c *gin.Context) (*CreateMultipleWallet, error) {
	var req CreateMultipleWallet
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func BindAddMembers(c *gin.Context) (*AddMembers, error) {
	var req AddMembers
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func BinAddWeight(c *gin.Context) (*AddWeight, error) {
	var req AddWeight
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func BinTxTransCation(c *gin.Context) (*TxTransCation, error) {
	var req TxTransCation
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
func BinNewTransCationNumber(c *gin.Context) (*NewTransCationNumber, error) {
	var req NewTransCationNumber
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func BindSignTxTransCation(c *gin.Context) (*SignTxTransCation, error) {
	var req SignTxTransCation
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func BindVerifyTransaction(c *gin.Context) (*VerifyTransaction, error) {
	var req VerifyTransaction
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func BindCancelTransaction(c *gin.Context) (*CancelTransaction, error) {
	var req CancelTransaction
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func BindTransactionList(c *gin.Context) (*TransactionList, error) {
	var req TransactionList
	if err := c.BindQuery(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func BindTxCompleted(c *gin.Context) (*TxCompleted, error) {
	var req TxCompleted
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func BinUpdateThreshold(c *gin.Context) (*UpdateThreshold, error) {
	var req UpdateThreshold
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func BinUpdateWeight(c *gin.Context) (*UpdateWeight, error) {
	var req UpdateWeight
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func BinGetUserInfo(c *gin.Context) (*GetUserInfo, error) {
	var req GetUserInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func BinVerifyTransactionBeReady(c *gin.Context) (*VerifyTransactionBeReady, error) {
	var req VerifyTransactionBeReady
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func BinTransactionHistory(c *gin.Context) (*TransactionHistory, error) {
	var req TransactionHistory
	if err := c.BindQuery(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func BinReplaceMemberAddressBeReady(c *gin.Context) (*ReplaceMemberAddress, error) {
	var req ReplaceMemberAddress
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
