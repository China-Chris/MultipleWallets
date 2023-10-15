package daos

import (
	"Multiplewallets/configs"
	"Multiplewallets/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMultipleSignatureWallet(t *testing.T) {
	configPath := "../config_test.yaml"
	configs.ParseConfig(configPath)
	// 初始化数据库连接
	InitMysql()
	// 创建一个 MultipleSignatureWallet 对象
	wallet := models.MultipleSignatureWallet{
		Address:   "0x6C8c3500A229Be8cdA5B7b166E5b4E40552c2aC3",
		Threshold: 2,
	}
	// 调用函数进行创建
	err := CreateMultipleSignatureWallet(wallet)
	assert.NoError(t, err)
}
