package service

import (
	"Multiplewallets/daos"
	"Multiplewallets/models"
	"Multiplewallets/request"
	"fmt"
	"time"
)

// CreateMultipleWallet 创建多签钱包的业务逻辑
func CreateMultipleWallet(req *request.CreateMultipleWallet) (string, error) {
	wallet := models.MultipleSignatureWallet{
		Address:    req.Address,
		Threshold:  req.Threshold,
		IsWeighted: req.IsWeighted,
	}
	// 创建钱包
	err := daos.CreateMultipleSignatureWallet(wallet)
	if err != nil {
		fmt.Println(err)
		return req.Address, err
	}
	return req.Address, nil
}

// AddMembers 添加门槛成员
func AddMembers(req *request.AddMembers) (bool, error) {
	//获取钱包
	wallet, err := daos.GetMultipleSignatureWalletByAddress(req.WalletAddress)
	if err != nil {
		fmt.Println(err)
	}
	memberSignature := models.MemberSignature{
		WalletId:      wallet.WalletId,
		Name:          req.Name,
		MemberAddress: req.Address,
	}
	err = daos.AddMemberSignature(memberSignature)
	if err != nil {
		fmt.Println(err)
	}
	return true, nil
}

// AddWeight 添加权重钱包
func AddWeight(req *request.AddWeight) (bool, error) {
	wallet, err := daos.GetMultipleSignatureWalletByAddress(req.WalletAddress)
	if err != nil {
		fmt.Println(err)
	}
	memberWeight := models.MemberWeight{
		WalletId:           wallet.WalletId,
		Name:               req.Name,
		MemberAddress:      req.Address,
		MemberWeightNumber: req.Weight,
	}
	err = daos.AddMemberWeight(memberWeight)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

// CreateTxTransCation 创建事务交易
func CreateTxTransCation(req *request.TxTransCation) (bool, error) {
	wallet, err := daos.GetMultipleSignatureWalletByAddress(req.WalletAddress)
	if err != nil {
		fmt.Println("未找到钱包", err)
		return false, err
	}
	fmt.Println("123")
	// 检查是否已经存在相同的交易
	existingTx, _ := daos.GetTransactionByWalletIDAndNonce(wallet.WalletId, req.Nonce)

	fmt.Println("321")
	if existingTx != nil {
		// 如果已存在相同的交易，更新它
		existingTx.Content = req.Content
		existingTx.IsModified = 2 //设置2为已经修改过
		err = daos.UpdateTransaction(*existingTx)
		if err != nil {
			fmt.Println(err)
			return false, err
		}
		return true, nil
	}
	txTransaction := models.Transaction{
		WalletId:        wallet.WalletId,
		Nonce:           req.Nonce,
		Threshold:       wallet.Threshold,
		TransactionType: req.TransactionType,
		Content:         req.Content,
	}
	err = daos.CreatTxTransCation(txTransaction)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

// GetNewTransCationNumber 获取最新事务序号
func GetNewTransCationNumber(req *request.NewTransCationNumber) (int, error) {
	wallet, err := daos.GetMultipleSignatureWalletByAddress(req.WalletAddress)
	if err != nil {
		fmt.Println(err)
	}
	nonce, err := daos.GetLatestNonceByWalletID(wallet.WalletId)
	if err != nil {
		return 0, nil
	}
	return nonce + 1, nil
}

func SignTxTransCation(req *request.SignTxTransCation) (bool, error) {
	wallet, err := daos.GetMultipleSignatureWalletByAddress(req.WalletAddress)
	if err != nil {
		fmt.Println(err)
	}
	transactionId, err := daos.FirstByWalletAddressAddNonce(wallet.WalletId, req.Nonce)
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	fmt.Println(req)
	signature := models.Signature{
		Address:       req.Address,
		WalletId:      wallet.WalletId,
		TransactionId: transactionId,
		SignatureData: req.SignatureData,
	}
	err = daos.AddSignTxTransCation(signature)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

// VerifyTransaction  TODO: 太长 需要拆分
func VerifyTransaction(req *request.VerifyTransaction) (list []struct {
	Address       string
	SignatureData string
}, err error) {

	fmt.Println(req)

	//先判断这个钱包是什么钱包
	wallet, err := daos.GetMultipleSignatureWalletByAddress(req.WalletAddress)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	fmt.Println(wallet)

	// 获取事务
	transaction, err := daos.GetTransactionAddressAddNonce(wallet.WalletId, req.Nonce)
	if err != nil {
		fmt.Println(err)
		return list, nil
	}
	signatureOrder, err := daos.GetSignatureOrderForTransaction(transaction.TransactionId)
	if err != nil {
		return nil, nil
	}
	fmt.Println(signatureOrder)
	// 将签名排序存储在 list 中，供调用方使用
	list = signatureOrder
	return list, nil
}

func CancelTransaction(req *request.CancelTransaction) (bool, error) {
	wallet, err := daos.GetMultipleSignatureWalletByAddress(req.WalletAddress)
	if err != nil {
		fmt.Println(err)
	}
	transactionId, err := daos.FirstByWalletAddressAddNonce(wallet.WalletId, req.Nonce)
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	// 调用dao撤销事务
	bool, err := daos.CancelTransaction(transactionId)
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	err = daos.DeleteSignaturesByAllAccountID(wallet.WalletId, transactionId)
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	return bool, nil
}

// TransactionLists 返回已签名地址列表和未签名地址列表
func TransactionLists(req *request.TransactionList) (signedAddresses []string, unsignedAddresses []string, err error) {
	// 先查询钱包地址
	wallet, err := daos.GetMultipleSignatureWalletByAddress(req.WalletAddress)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("已经查询钱包地址：", wallet)
	// 查询钱包对应的交易

	fmt.Println(req.Nonce)
	transaction, err := daos.GetTransactionAddressAddNonce(wallet.WalletId, req.Nonce)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(2)
	count, _ := daos.CountNumberOfTransactionSignatures(transaction.TransactionId)
	if err != nil {
		fmt.Println(count)
		return nil, nil, err
	}
	fmt.Println("count", count)
	if count == 0 {
		if wallet.IsWeighted == 1 {
			// 获取门限钱包的成员列表
			members, err := daos.GetThresholdWalletMembers(wallet.WalletId)
			if err != nil {
				fmt.Println("报错", err)
				return nil, nil, err
			}
			fmt.Println("members", members)
			// 找出未签名的成员
			unsignedAddresses := make([]string, 0)
			for _, member := range members {
				unsignedAddresses = append(unsignedAddresses, member.MemberAddress)
			}
			return nil, unsignedAddresses, nil
		} else {
			fmt.Println("_________获取权重钱包成员列表_______")
			// 获取权重钱包的成员列表
			members, err := daos.GetWeightedWalletMembers(wallet.WalletId)
			if err != nil {
				fmt.Println("报错", err)
				return nil, nil, err
			}

			fmt.Println("members", members)
			// 找出未签名的成员
			unsignedAddresses := make([]string, 0)
			for _, member := range members {
				fmt.Println(member.MemberAddress)
				unsignedAddresses = append(unsignedAddresses, member.MemberAddress)
			}

			fmt.Println("unsignedAddresses", unsignedAddresses)
			return nil, unsignedAddresses, nil
		}
	}
	fmt.Println("没有签名")
	// 从 signature 表中获取当前已经签名的用户地址列表
	signatureAddresses, err := daos.GetSignatureAddressesByTransactionID(transaction.TransactionId)
	if err != nil {
		fmt.Println("123")
		return nil, nil, nil
	}
	// 使用 map 存储签名地址以提高查找性能
	signatureMap := make(map[string]struct{})
	for _, address := range signatureAddresses {
		signatureMap[address] = struct{}{}
	}
	fmt.Println("345")
	fmt.Println(wallet.IsWeighted)
	if wallet.IsWeighted != 1 { // 是权重钱包
		// 获取权重钱包的成员列表
		members, err := daos.GetWeightedWalletMembers(wallet.WalletId)
		if err != nil {
			return nil, nil, err
		}
		// 找出未签名的成员
		for _, member := range members {
			// 检查成员是否在签名地址列表中
			if _, ok := signatureMap[member.MemberAddress]; !ok {
				unsignedAddresses = append(unsignedAddresses, member.MemberAddress)
			} else {
				signedAddresses = append(signedAddresses, member.MemberAddress)
			}
		}
		return signedAddresses, unsignedAddresses, nil
	} else { // 是门限钱包
		// 获取门限钱包的成员列表
		members, err := daos.GetThresholdWalletMembers(wallet.WalletId)
		if err != nil {
			return nil, nil, err
		}
		fmt.Println(members)
		// 计算已签名成员数量
		signedMemberCount := 0
		for _, member := range members {
			// 检查成员是否在签名地址列表中
			if _, ok := signatureMap[member.MemberAddress]; ok {
				signedMemberCount++
				signedAddresses = append(signedAddresses, member.MemberAddress)
			} else {
				unsignedAddresses = append(unsignedAddresses, member.MemberAddress)
			}
		}
		// 检查门限值是否足够
		if signedMemberCount >= wallet.Threshold {
			return signedAddresses, unsignedAddresses, nil
		} else {
			return signedAddresses, unsignedAddresses, nil
		}
	}
}

func TxCompleted(req *request.TxCompleted) (bool, error) {
	tx, err := daos.StartDatabaseTransaction()
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	defer tx.Rollback()

	wallet, err := daos.GetMultipleSignatureWalletByAddress(req.WalletAddress)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	transactionID, err := daos.FirstByWalletAddressAddNonce(wallet.WalletId, req.Nonce)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	err = daos.UpdateTransactionHash(transactionID, req.Hash)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	//将当前钱包的成员保存到TransactionMemberInfo表中
	switch wallet.IsWeighted {
	case 1:
		// 门限钱包
		membersThreshold, err := daos.GetThresholdWalletMembers(wallet.WalletId)
		if err != nil {
			fmt.Println(err)
			return false, err
		}
		for _, memberThreshold := range membersThreshold {
			memberInfo := models.TransactionMemberInfo{
				TransactionID: transactionID,
				MemberName:    memberThreshold.Name,
				MemberAddress: memberThreshold.MemberAddress,
				CreatedAt:     time.Now(),
			}
			err = daos.CreateTransactionMemberInfo(memberInfo)
			if err != nil {
				fmt.Println(err)
				return false, err
			}
		}

	case 2:
		// 权重钱包
		membersWeighted, err := daos.GetWeightedWalletMembers(wallet.WalletId)
		if err != nil {
			fmt.Println(err)
			return false, err
		}
		for _, memberWeighted := range membersWeighted {
			memberInfo := models.TransactionMemberInfo{
				TransactionID: transactionID,
				MemberName:    memberWeighted.Name,
				MemberAddress: memberWeighted.MemberAddress,
				CreatedAt:     time.Now(),
			}
			err = daos.CreateTransactionMemberInfo(memberInfo)
			if err != nil {
				fmt.Println(err)
				return false, err
			}
		}
	default:
		return false, nil
	}
	tx.Commit()
	return true, nil
}

// UpdateThreshold
func UpdateThreshold(req *request.UpdateThreshold) (bool, error) {
	wallet, err := daos.GetMultipleSignatureWalletByAddress(req.WalletAddress)
	if err != nil {
		fmt.Println(err)
	}
	err = daos.UpdateThreshold(wallet.WalletId, req.Threshold)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// UpdateWeight
func UpdateWeight(req *request.UpdateWeight) (bool, error) {
	wallet, err := daos.GetMultipleSignatureWalletByAddress(req.WalletAddress)
	if err != nil {
		fmt.Println(err)
	}
	//修改单个用户的权重
	err = daos.UpdateMemberWeight(wallet.WalletId, req.Address, req.Weight)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetUserInfo(req *request.GetUserInfo) (string, string, error) {
	//先判断这个钱包是什么钱包
	wallet, err := daos.GetMultipleSignatureWalletByAddress(req.WalletAddress)
	if err != nil {
		fmt.Println(err)
		return "", "", nil
	}
	if wallet.IsWeighted == 1 { //是门限值钱包
		//从门限成员处获取
		wallet, err := daos.GetMemberSignature(req.Address)
		if err != nil {
			fmt.Println(err)
			return "", "", nil
		}
		return wallet.Name, wallet.Img, nil
	} else { // 如果是权重钱包

		wallet, err := daos.GetMemberWeight(req.Address)
		if err != nil {
			fmt.Println(err)
			return "", "", nil
		}
		return wallet.Name, wallet.Img, nil
	}
	return "", "", nil
}

func VerifyTransactionBeReady(req *request.VerifyTransactionBeReady) (bool, error) {
	fmt.Println(req)
	//先判断这个钱包是什么钱包
	wallet, err := daos.GetMultipleSignatureWalletByAddress(req.WalletAddress)
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	fmt.Println(wallet)
	// 获取事务
	transaction, err := daos.GetTransactionAddressAddNonce(wallet.WalletId, req.Nonce)
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	fmt.Println(transaction)

	if wallet.IsWeighted == 1 { //是门限值钱包
		//统计钱包该事务的签名数量
		count, err := daos.CountNumberOfTransactionSignatures(wallet.WalletId)
		if err != nil {
			return false, nil
		}
		if count >= wallet.Threshold { //当 签名数量大于门限值  返回对该笔事务的签名排序

			return true, nil
		}

	} else { //是权重钱包  统计该事务签名的地址在权重钱包成员中的权重总和
		fmt.Println("123")
		// 获取权重钱包的成员列表
		members, err := daos.GetWeightedWalletMembers(wallet.WalletId)
		if err != nil {
			return false, nil
		}
		fmt.Println("members", members)

		// 初始化用于存储成员权重的映射
		memberWeights := make(map[string]int)

		// 统计成员在签名数据中出现的次数
		for _, member := range members {
			// 在签名数据中搜索成员地址
			count, err := daos.CountSignatureDataForMember(member.MemberAddress, transaction.TransactionId, wallet.WalletId)
			if err != nil {
				fmt.Println("12345")
				return false, err
			}
			// 存储成员的权重
			memberWeights[member.MemberAddress] = count
		}

		// 计算成员权重的综合
		totalWeight := 0 //权重和
		for _, member := range members {
			totalWeight += memberWeights[member.MemberAddress]
		}
		fmt.Println(totalWeight)

		//比较成员综合权重值和钱包定义的权重值
		if totalWeight >= wallet.Threshold { //成员综合权重值大于钱包权重值
			return true, nil
		}
	}
	return false, nil
}

func TransactionHistory(req *request.TransactionHistory) ([]models.Transaction, error) {
	//先判断这个钱包是什么钱包
	wallet, err := daos.GetMultipleSignatureWalletByAddress(req.WalletAddress)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	Tran, err := daos.GetTransactionAddress(wallet.WalletId)
	if err != nil {
		return nil, err
	}
	return Tran, nil
}

func DeleteMember(req *request.DeleteMember) (bool, error) {
	wallet, err := daos.GetMultipleSignatureWalletByAddress(req.WalletAddress)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	switch wallet.IsWeighted {
	case 1:
		// 门限钱包
		err = daos.DeleteMember(wallet.WalletId, req.Address)
		if err != nil {
			return false, nil
		}
	case 2:
		// 权重钱包
		err = daos.DeleteMemberWeight(wallet.WalletId, req.Address)
		if err != nil {
			return false, nil
		}
	default:
		return false, nil
	}
	//删除账户所对于钱包的签名
	err = daos.DeleteSignaturesByAccountID(req.Address, wallet.WalletId)
	if err != nil {
		return false, err
	}

	return true, nil
}

func ReplaceMemberAddress(req *request.ReplaceMemberAddress) (bool, error) {
	wallet, err := daos.GetMultipleSignatureWalletByAddress(req.WalletAddress)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	fmt.Println(wallet.IsWeighted)
	switch wallet.IsWeighted {
	case 1:
		// 门限钱包
		err := daos.ReplaceThresholdWalletMemberAddress(wallet.WalletId, req.Address, req.NewAddress)
		if err != nil {
			fmt.Println(err)
			return false, err
		}

	case 2:
		// 权重钱包
		// 在这里添加处理权重钱包的逻辑
		err := daos.ReplaceWeightedWalletMemberAddress(wallet.WalletId, req.Address, req.NewAddress)
		if err != nil {
			fmt.Println(err)
			return true, err
		}
	default:
		return false, nil
	}
	return true, nil
}
