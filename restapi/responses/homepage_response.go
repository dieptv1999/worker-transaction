package responses

import "github.com/OpenStars/EtcdBackendService/RankStorageService/models"

type HomePageData struct {
	ItemsBanner           []models.Item `json:"items_banner"`
	ItemsView             []models.Item `json:"items_view"`
	AuctionsView          []models.Item `json:"auctions_view"`
	TotalItem             int64         `json:"total_item"`
	TotalAuction          int64         `json:"total_auction"`
	TotalCollection       int64         `json:"total_collection"`
	TotalFeaturedCreators int64         `json:"total_featured_creators"`
}
type HomePageResponse struct {
	Data  *HomePageData `json:"data"`
	Error *Err          `json:"error"`
}
