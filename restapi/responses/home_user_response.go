package responses

import "github.com/OpenStars/EtcdBackendService/RankStorageService/models"

type HomeUserData struct {
	ItemsView       []models.Item `json:"items_view"`
	TotalItem       int64         `json:"total_item"`
	TotalCollection int64         `json:"total_collection"`
}
type HomeUserResponse struct {
	Data  *HomeUserData `json:"data"`
	Error *Err          `json:"error"`
}
