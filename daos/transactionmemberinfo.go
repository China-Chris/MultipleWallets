package daos

import "Multiplewallets/models"

func CreateTransactionMemberInfo(wallet models.TransactionMemberInfo) (err error) {
	return db.Create(&wallet).Error
}
