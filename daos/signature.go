package daos

import (
	"Multiplewallets/models"
)

// AddSignTxTransCation
func AddSignTxTransCation(wallet models.Signature) (err error) {
	return db.Create(&wallet).Error
}

// GetTransactionByWalletIDAndNonce 根据钱包ID和事务序号获取事务
func GetTransactionSignatures(transactionId int) (*models.Signature, error) {
	var signature models.Signature
	if err := db.Where("transaction_id = ?").First(&signature).Error; err != nil {
		return nil, err
	}
	return &signature, nil
}

// 统计签名数量
func CountNumberOfTransactionSignatures(transactionId int) (int, error) {
	var count int64
	err := db.Model(&models.Signature{}).
		Where("transaction_id = ?", transactionId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

// GetSignatureAddressesByTransactionID 获取与给定交易ID关联的签名地址列表
func GetSignatureAddressesByTransactionID(transactionID int) ([]string, error) {
	// 查询与给定交易ID关联的签名地址列表
	var signatureAddresses []string
	err := db.Model(&models.Signature{}).
		Where("transaction_id = ?", transactionID).
		Pluck("address", &signatureAddresses).Error
	if err != nil {
		return nil, err // 返回错误，而不是返回 nil, nil
	}
	return signatureAddresses, nil
}

// DeleteSignaturesByAccountID 根据账户ID删除所有相关的签名
func DeleteSignaturesByAccountID(address string, wallet int) error {
	// 使用事务来删除与给定账户ID关联的所有签名
	if err := db.Where("address = ? AND wallet_id = ?", address, wallet).Delete(&models.Signature{}).Error; err != nil {
		return err
	}
	return nil
}

// DeleteSignaturesByAllAccountID 根据钱包id和交易id删除所有相关的签名
func DeleteSignaturesByAllAccountID(walletID int, transactionID int) error {
	// 使用事务来删除与给定地址和钱包ID关联的所有签名
	if err := db.Where("wallet_id = ? AND transaction_id = ?", walletID, transactionID).Delete(&models.Signature{}).Error; err != nil {
		return err
	}
	return nil
}
