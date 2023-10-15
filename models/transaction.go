package models

import (
	"time"
)

// Transaction represents the data structure for a transaction in a multiple signature wallet.
type Transaction struct {
	TransactionId   int       `gorm:"primary_key";autoIncrement`
	WalletId        int       // 外键，关联到对应的钱包
	Nonce           int       // Transaction  交易序号
	TransactionType string    // 交易类型（例如：转账、合约）
	Content         string    // 交易内容的 JSON 格式
	IsModified      int       // 表示交易是否被修改过
	Hash            string    // 表示交易是否已执行+
	Threshold       int       // 发生交易时候的门限值
	ExecutedAt      time.Time // 交易执行的时间戳
	CreatedAt       time.Time // 创建时间

	// 添加索引
	IndexWalletId int    `gorm:"index"` // 为外键 WalletId 添加索引
	IndexNonce    int    `gorm:"index"` // 为 Nonce 添加索引
	IndexHash     string `gorm:"index"` // 为 Hash 添加索引
}

// TableName returns the corresponding database table name for this struct.
func (t Transaction) TableName() string {
	return "transaction"
}
