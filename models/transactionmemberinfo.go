package models

import "time"

// Transaction represents the data structure for a transaction in a multiple signature wallet.
type TransactionMemberInfo struct {
	TransactionMemberInfoId int       `gorm:"primary_key";autoIncrement`
	TransactionID           int       // 外键，关联到对应的事务ID
	MemberName              string    // 成员名称
	MemberAddress           string    // 成员地址
	CreatedAt               time.Time // 创建时间

	// 添加索引
	IndexTransactionID int    `gorm:"index"` // 为外键 TransactionID 添加索引
	IndexMemberName    string `gorm:"index"` // 为 MemberName 添加索引
	IndexMemberAddress string `gorm:"index"` // 为 MemberAddress 添加索引
}

// TableName returns the corresponding database table name for this struct.
func (t TransactionMemberInfo) TableName() string {
	return "transaction_member_info"
}
