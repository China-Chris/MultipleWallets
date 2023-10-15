package models

import "time"

//这张表用于存储多签钱包的基本信息。每一行记录代表一个多签钱包，包括其地址、门限值、是否为权重值钱包等信息。这个表可以让您在数据库中保存每个多签钱包的特定属性。
//This table is used to store basic information about multi signature wallets. Each row of records represents a multi signature wallet, including its address, threshold value, whether it is a weight value wallet, and other information. This table allows you to save specific attributes for each multi signed wallet in the database.

// MultipleSignatureWallet represents the data structure for a multiple signature wallet.
type MultipleSignatureWallet struct {
	WalletId   int       `gorm:"primary_key"; autoIncrement`
	Address    string    //钱包地址
	Threshold  int       //多签签名的门限值
	IsWeighted int       //标识是否为权重钱包    1是门限钱包  2是权重钱包
	CreatedAt  time.Time //创建时间
}

// TableName returns the corresponding database table name for this struct.
func (m MultipleSignatureWallet) TableName() string {
	return "multiple_signature_wallet"
}
