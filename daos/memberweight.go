package daos

import "Multiplewallets/models"

// AddMemberWeight 添加权重
func AddMemberWeight(wallet models.MemberWeight) (err error) {
	return db.Create(&wallet).Error
}

// GetWeightedWalletMembers 获取权重钱包的成员列表
func GetWeightedWalletMembers(walletID int) ([]models.MemberWeight, error) {
	var members []models.MemberWeight
	// 查询指定钱包ID的成员列表
	if err := db.Where("wallet_id = ?", walletID).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

// UpdateMemberWeight 更新成员的权重
func UpdateMemberWeight(walletID int, memberAddress string, memberWeightNumber int) error {
	// 使用 Update 方法直接更新成员的权重
	if err := db.Model(models.MemberWeight{}).
		Where("wallet_id = ? AND member_address = ?", walletID, memberAddress).
		Update("member_weight_number", memberWeightNumber).Error; err != nil {
		return err
	}
	return nil
}

func GetMemberWeight(address string) (wallet models.MemberSignature, err error) {
	wallet = models.MemberSignature{}
	err = db.Where("member_address = ?", address).First(&wallet).Error
	return wallet, err
}

// DeleteMemberWeight 删除成员
func DeleteMemberWeight(walletID int, memberAddress string) error {
	// 使用 Delete 方法删除指定钱包ID和成员地址的成员
	if err := db.Where("wallet_id = ? AND member_address = ?", walletID, memberAddress).Delete(models.MemberWeight{}).Error; err != nil {
		return err
	}
	return nil
}

// ReplaceThresholdWalletMemberAddress
func ReplaceWeightedWalletMemberAddress(walletID int, oldMemberAddress string, newMemberAddress string) error {
	// 使用 Update 方法来替换门槛钱包成员的地址
	if err := db.Model(models.MemberWeight{}).
		Where("wallet_id = ? AND member_address = ?", walletID, oldMemberAddress).
		Update("member_address", newMemberAddress).Error; err != nil {
		return err
	}
	return nil
}
