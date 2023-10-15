package models

import (
	"time"
)

// MemberSignature represents the data structure for a member's signature on a transaction in a multiple signature wallet.
type MemberSignature struct {
	MemberSignatureId int `gorm:"primary_key";autoIncrement`
	WalletId          int // 外键，关联到对应的WalletId
	Name              string
	Img               string
	MemberAddress     string    // 成员地址
	CreatedAt         time.Time // 创建时间

	// 添加索引
	IndexWalletID      int    `gorm:"index"` // 为外键添加索引
	IndexMemberAddress string `gorm:"index"` // 为成员地址添加索引
}

// TableName returns the corresponding database table name for this struct.
func (m MemberSignature) TableName() string {
	return "member_signature"
}
