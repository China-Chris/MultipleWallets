package response

type CreateMultipleWalletResponse struct {
	Address string `json:"address"`
}

type AddMembersResponse struct {
	Bool bool `json:"bool"`
}
type AddWeightResponse struct {
	Bool bool `json:"bool"`
}

type CreateTxTransCationResponse struct {
	Bool bool `json:"bool"`
}

type NewTransCationNumberResponse struct {
	Nonce int `json:"nonce"`
}

type SignTxTransCationResponse struct {
	Bool bool `json:"bool"`
}

type CancelTransactionResponse struct {
	Bool bool `json:"bool"`
}

type TransactionListResponse struct {
	SignedAddresses   []string `json:"signed_addresses"`
	UnsignedAddresses []string `json:"unsigned_addresses"`
}

type TxCompletedResponse struct {
	Bool bool `json:"bool"`
}

type UpdateThresholdResponse struct {
	Bool bool `json:"bool"`
}

type GetUserInfoResponse struct {
	Name string `json:"name"`
	Img  string `json:"img"`
}

type VerifyTransactionBeReady struct {
	Bool bool `json:"bool"`
}
