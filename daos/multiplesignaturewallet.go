package daos

import "Multiplewallets/models"

func GetMultipleSignatureWalletByAddress(address string) (wallet models.MultipleSignatureWallet, err error) {
	wallet = models.MultipleSignatureWallet{}
	err = db.Where("address = ?", address).First(&wallet).Error
	return wallet, err
}

func CreateMultipleSignatureWallet(wallet models.MultipleSignatureWallet) (err error) {
	return db.Create(&wallet).Error
}

func UpdateThreshold(walletID int, newThreshold int) (err error) {
	err = db.Model(models.MultipleSignatureWallet{}).
		Where("wallet_id = ?", walletID).
		Update("threshold", newThreshold).
		Error
	return err
}
