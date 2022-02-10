package responses

import "math/big"

type DataMarketplaceResponse struct {
	DataApprove *DataApprove `json:"data_approve"`
	Data        *Data        `json:"data"`
	Error       *Err         `json:"error"`
}

type Data struct {
	From     string   `json:"from"`
	To       string   `json:"to"`
	Gas      uint64   `json:"gas"`
	GasPrice *big.Int `json:"gas_price"`
	Value    *big.Int `json:"value"`
	Data     string   `json:"data"`
}

type DataResponse struct {
	Data  *Data `json:"data"`
	Error *Err  `json:"error"`
}

type DataApproveResponse struct {
	Data  *DataApprove `json:"data"`
	Error *Err         `json:"error"`
}

type DataApprove struct {
	From     string   `json:"from"`
	To       string   `json:"to"`
	Gas      uint64   `json:"gas"`
	GasPrice *big.Int `json:"gas_price"`
	Value    *big.Int `json:"value"`
	Data     string   `json:"data"`
}

type DataFee struct {
	FeeBuy    float64 `json:"fee_buy"`
	FeeSell   float64 `json:"fee_sell"`
	Royalties float64 `json:"royalties"`
}

type DataAuctionFee struct {
	FeeAuction float64 `json:"fee_auction"`
	Royalties  float64 `json:"royalties"`
}

type DataFeeResponse struct {
	Data  *DataFee `json:"data"`
	Error *Err     `json:"error"`
}

type DataAuctionFeeResponse struct {
	Data  *DataAuctionFee `json:"data"`
	Error *Err            `json:"error"`
}
