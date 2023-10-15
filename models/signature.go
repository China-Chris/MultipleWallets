package models

import "time"

type Signature struct {
	SignatureId   int       `gorm:"primary_key;autoIncrement"`
	Address       string    // 签名用户地址
	WalletId      int       // 外键，关联到对应的WalletId
	TransactionId int       // 外键，关联到对应的TransactionId
	SignatureData string    // 数字签名数据
	CreatedAt     time.Time // 创建时间

	// 添加索引
	IndexAddress       string `gorm:"index"` // 为地址添加索引
	IndexWalletId      int    `gorm:"index"` // 为外键 WalletId 添加索引
	IndexTransactionId int    `gorm:"index"` // 为外键 TransactionId 添加索引
}

// TableName returns the corresponding database table name for this struct.
func (m Signature) TableName() string {
	return "signature"
}
