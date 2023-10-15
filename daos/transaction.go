package daos

import (
	"Multiplewallets/models"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// CreatTxTransCation 创建事务交易
func CreatTxTransCation(wallet models.Transaction) (err error) {
	return db.Create(&wallet).Error
}

// GetTransactionByWalletIDAndNonce 根据钱包ID和事务序号获取事务
func GetTransactionByWalletIDAndNonce(walletID int, nonce int) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := db.Where("wallet_id = ? AND nonce = ?", walletID, nonce).First(&transaction).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func GetLatestNonceByWalletID(walletID int) (nonce int, err error) {
	var transaction models.Transaction
	err = db.
		Where("wallet_id = ?", walletID).
		Order("nonce desc").
		First(&transaction).
		Error

	if err != nil {
		return 0, err
	}

	return transaction.Nonce, nil
}

func FirstByWalletAddressAddNonce(walletID int, nonce int) (int, error) {
	var transaction models.Transaction
	err := db.
		Where("wallet_id = ? AND nonce = ?", walletID, nonce).
		First(&transaction).
		Error

	if err != nil {
		return 0, err
	}
	return transaction.TransactionId, nil
}

func GetTransactionAddressAddNonce(walletID int, nonce int) (transaction models.Transaction, err error) {
	transaction = models.Transaction{}
	err = db.Where("wallet_id = ? AND nonce = ?", walletID, nonce).
		First(&transaction).Error
	return transaction, err
}

func GetTransactionAddress(walletID int) (transactions []models.Transaction, err error) {
	err = db.Where("wallet_id = ?", walletID).
		Find(&transactions).Error
	return transactions, err
}

func GetSignatureOrderForTransaction(transactionId int) (signatureDataWithAddress []struct {
	Address       string
	SignatureData string
}, err error) {
	var signatures []models.Signature
	// 查询特定事务的所有签名记录，并按照地址升序排序
	if err := db.Where("transaction_id = ?", transactionId).Order("address ASC").Find(&signatures).Error; err != nil {
		return nil, err
	}
	// 初始化用于存储签名数据和签名地址的切片
	signatureDataWithAddress = make([]struct {
		Address       string
		SignatureData string
	}, len(signatures))

	// 提取签名数据和签名地址并存储在 signatureDataWithAddress 中
	for i, sig := range signatures {
		signatureDataWithAddress[i] = struct {
			Address       string
			SignatureData string
		}{
			Address:       sig.Address,
			SignatureData: sig.SignatureData,
		}
	}
	return signatureDataWithAddress, nil
}

// CountSignatureDataForMember 统计指定成员在指定事务中的签名权重
func CountSignatureDataForMember(memberAddress string, transactionID int, walletId int) (number int, err error) {
	var signatures models.Signature
	var memberWeight models.MemberWeight
	// 查询指定成员地址在指定事务中的签名数量
	err = db.
		Where("address = ? AND transaction_id = ?", memberAddress, transactionID).
		Find(&signatures).
		Error
	if err != nil {
		return 0, err
	}
	fmt.Println("address", signatures.Address, walletId)
	err = db.Where("member_address = ? AND wallet_id= ?", signatures.Address, walletId).
		Find(&memberWeight).Error
	return memberWeight.MemberWeightNumber, nil
}

// CancelTransaction 根据事务ID取消事务的DAO
func CancelTransaction(transactionID int) (bool, error) {
	// 在数据库中查找特定事务
	var transaction models.Transaction
	err := db.Where("transaction_id = ?", transactionID).First(&transaction).Error
	if err != nil {
		return false, err // 如果未找到事务或发生错误，返回错误
	}
	// 执行取消操作，例如将事务状态设置为已取消
	transaction.IsModified = 1
	// 更新事务状态
	err = db.Save(&transaction).Error
	if err != nil {
		return false, err // 如果更新时发生错误，返回错误
	}
	// 如果取消操作成功，返回 nil 表示没有错误
	return true, nil
}

// UpdateTransactionHash 更新 Transaction 表中的 Hash 字段
func UpdateTransactionHash(transactionID int, newHash string) error {
	// 在事务中执行更新操作，以确保数据的一致性
	return db.Transaction(func(tx *gorm.DB) error {
		// 检查是否存在具有给定 TransactionId 的记录
		var existingTransaction models.Transaction
		if err := tx.Where("transaction_id = ?", transactionID).First(&existingTransaction).Error; err != nil {
			return err
		}

		// 更新 Hash 字段
		existingTransaction.Hash = newHash
		existingTransaction.ExecutedAt = time.Now()

		// 保存更新后的记录
		if err := tx.Save(&existingTransaction).Error; err != nil {
			return err
		}

		return nil
	})
}

func UpdateTransaction(transaction models.Transaction) (err error) {
	err = db.Save(&transaction).Error
	return err
}
