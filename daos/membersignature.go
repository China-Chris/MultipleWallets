package daos

import "Multiplewallets/models"

// AddMemberSignature 添加门槛成员
func AddMemberSignature(wallet models.MemberSignature) (err error) {
	return db.Create(&wallet).Error
}

func GetMemberSignature(address string) (wallet models.MemberSignature, err error) {
	wallet = models.MemberSignature{}
	err = db.Where("member_address = ?", address).First(&wallet).Error
	return wallet, err
}

func GetThresholdWalletMembers(walletID int) ([]models.MemberSignature, error) {
	var members []models.MemberSignature
	// 查询指定钱包ID的成员列表
	if err := db.Where("wallet_id = ?", walletID).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

// DeleteMember 删除成员
func DeleteMember(walletID int, memberAddress string) error {
	// 使用 Delete 方法删除指定钱包ID和成员地址的成员
	if err := db.Where("wallet_id = ? AND member_address = ?", walletID, memberAddress).Delete(models.MemberSignature{}).Error; err != nil {
		return err
	}
	return nil
}

// ReplaceThresholdWalletMemberAddress
func ReplaceThresholdWalletMemberAddress(walletID int, oldMemberAddress string, newMemberAddress string) error {
	// 使用 Update 方法来替换门槛钱包成员的地址
	if err := db.Model(models.MemberSignature{}).
		Where("wallet_id = ? AND member_address = ?", walletID, oldMemberAddress).
		Update("member_address", newMemberAddress).Error; err != nil {
		return err
	}
	return nil
}
