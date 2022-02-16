package requests

type WithdrawalRequest struct {
	Coin      string `json:"coin" xml:"coin"`
	Chain     string `json:"chain" xml:"chain"`
	Amount    string `json:"amount" xml:"amount"`
	ToAddress string `json:"to_address" xml:"to_address"`
}
