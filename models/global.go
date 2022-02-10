package models

import (
	"log"
	"sync"
	"worker-transaction/consts"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
)

var (
	mapBigset map[string]*MyStringBSService
	onceBS    sync.Once
)

type MyStringBSService struct {
	StringBigsetService.StringBigsetServiceIf
}

func InitBigSetIf() {
	log.Println("=== @InitBigSetIf ===")
	//onceBS.Do(func() {
	//	mapBigset = map[string]*MyStringBSService{}
	//
	//	mapBigset[consts.BS_NFT_ORIGIN_SERVICE] = &MyStringBSService{
	//		bigsetservice.GetBigSet(appconfig.Config.NftMarketplaceSid,
	//			appconfig.Config.NftMarketplaceHost, appconfig.Config.NftMarketplacePort)}
	//
	//	mapBigset[consts.BS_NFT_ITEM_SERVICE] = &MyStringBSService{
	//		bigsetservice.GetBigSet(appconfig.Config.NftItemSid,
	//			appconfig.Config.NftItemHost, appconfig.Config.NftItemPort)}
	//
	//	mapBigset[consts.BS_NFT_USER_SERVICE] = &MyStringBSService{
	//		bigsetservice.GetBigSet(appconfig.Config.NftUserSid,
	//			appconfig.Config.NftUserHost, appconfig.Config.NftUserPort)}
	//
	//	mapBigset[consts.BS_NFT_EVENT_SERVICE] = &MyStringBSService{
	//		bigsetservice.GetBigSet(appconfig.Config.NftEventSid,
	//			appconfig.Config.NftEventHost, appconfig.Config.NftEventPort)}
	//
	//	mapBigset[consts.BS_NFT_ITEM_MAPPING_SERVICE] = &MyStringBSService{
	//		bigsetservice.GetBigSet(appconfig.Config.NftItemMappingSid,
	//			appconfig.Config.NftItemMappingHost, appconfig.Config.NftItemMappingPort)}
	//})
}

func GetBsNftOriginal() *MyStringBSService {
	return mapBigset[consts.BS_NFT_ORIGIN_SERVICE]
}

func GetBsNftItem() *MyStringBSService {
	return mapBigset[consts.BS_NFT_ITEM_SERVICE]
}

func GetBsNftMappingItem() *MyStringBSService {
	return mapBigset[consts.BS_NFT_ITEM_MAPPING_SERVICE]
}

func GetBsNftUser() *MyStringBSService {
	return mapBigset[consts.BS_NFT_USER_SERVICE]
}

func GetBsNftEvent() *MyStringBSService {
	return mapBigset[consts.BS_NFT_EVENT_SERVICE]
}
