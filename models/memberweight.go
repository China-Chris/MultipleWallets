package models

import (
	"time"
)

type MemberWeight struct {
	MemberWeightId     int       `gorm:"primary_key";autoIncrement`
	WalletId           int       // 外键，关联到对应的WalletId
	Name               string    // 名称
	MemberAddress      string    // 成员地址
	MemberWeightNumber int       // 成员权重
	CreatedAt          time.Time // 创建时间

	// 添加索引
	IndexWalletID      int    `gorm:"index"` // 为外键添加索引
	IndexMemberAddress string `gorm:"index"` // 为成员地址添加索引
}

// TableName returns the corresponding database table name for this struct.
func (m MemberWeight) TableName() string {
	return "member_weight"
}
