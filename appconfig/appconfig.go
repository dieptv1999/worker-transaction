package appconfig

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"worker-transaction/helps"
)

//go:generate easytags $GOFILE json,xml

type AppConfig struct {
	EtcdServerEndpoints []string `json:"etcd_server_endpoints" xml:"etcd_server_endpoints"`
	BnbRPC              string   `json:"bnb_rpc" xml:"bnb_rpc"`
	BnbRPCS             []string `json:"bnb_rpcs" xml:"bnb_rpcs"`
	PoolCap             int      `json:"pool_cap" xml:"pool_cap"`
	WkPoolBsc           int      `json:"wk_pool_bsc" xml:"wk_pool_bsc"`
	NumRetry            int      `json:"num_retry" xml:"num_retry"`

	TimeCache    int64 `json:"time_cache" xml:"time_cache"`
	IsSessionKey bool  `json:"is_session_key" xml:"is_sessionkey"`

	Duration int64 `json:"duration" xml:"duration"`

	NFTThriftHost string `json:"nft_thrift_host" xml:"nft_thrift_host"`
	NFTThriftPort string `json:"nft_thrift_port" xml:"nft_thrift_port"`

	RedisHost     string `json:"redis_host" xml:"redis_host"`
	RedisPassword string `json:"redis_password" xml:"redis_password"`
	RedisPoolSize int    `json:"redis_pool_size" xml:"redis_pool_size"`

	SessionServiceSid      string `json:"session_service_sid" xml:"session_service_sid"`
	SessionServiceHost     string `json:"session_service_host" xml:"session_service_host"`
	SessionServicePort     string `json:"session_service_port" xml:"session_service_port"`
	SessionServiceProtocol string `json:"session_service_protocol" xml:"session_service_protocol"`

	Pubkey2UidSid  string `json:"pubkey2_uid_service_sid" xml:"pubkey2_uid_service_sid"`
	Pubkey2UidHost string `json:"pubkey2_uid_host" xml:"pubkey2_uid_host"`
	Pubkey2UidPort string `json:"pubkey2_uid_port" xml:"pubkey2_uid_port"`

	KvCounterSid  string `json:"kv_counter_sid" xml:"kv_counter_sid"`
	KvCounterHost string `json:"kv_counter_host" xml:"kv_counter_host"`
	KvCounterPort string `json:"kv_counter_port" xml:"kv_counter_port"`

	NftMarketplaceSid  string `json:"nft_marketplace_sid" xml:"nft_marketplace_sid"`
	NftMarketplaceHost string `json:"nft_marketplace_host" xml:"nft_marketplace_host"`
	NftMarketplacePort string `json:"nft_marketplace_port" xml:"nft_marketplace_port"`

	NftItemSid  string `json:"nft_item_sid" xml:"nft_item_sid"`
	NftItemHost string `json:"nft_item_host" xml:"nft_item_host"`
	NftItemPort string `json:"nft_item_port" xml:"nft_item_port"`

	NftItemMappingSid  string `json:"nft_item_mapping_sid" xml:"nft_item_mapping_sid"`
	NftItemMappingHost string `json:"nft_item_mapping_host" xml:"nft_item_mapping_host"`
	NftItemMappingPort string `json:"nft_item_mapping_port" xml:"nft_item_mapping_port"`

	NftUserSid  string `json:"nft_user_sid" xml:"nft_user_sid"`
	NftUserHost string `json:"nft_user_host" xml:"nft_user_host"`
	NftUserPort string `json:"nft_user_port" xml:"nft_user_port"`

	NftEventSid  string `json:"nft_event_sid" xml:"nft_event_sid"`
	NftEventHost string `json:"nft_event_host" xml:"nft_event_host"`
	NftEventPort string `json:"nft_event_port" xml:"nft_event_port"`

	HttpHost string `json:"httphost" xml:"httphost"`
	HttpPort int    `json:"httpport" xml:"httpport"`
	Runmode  string `json:"run_mode" xml:"run_mode"`

	MailGunServiceHost     string `json:"mail_gun_service_host" xml:"mail_gun_service_host"`
	MailGunServicePort     string `json:"mail_gun_service_port" xml:"mail_gun_service_port"`
	MailGunServiceSid      string `json:"mail_gun_service_sid" xml:"mail_gun_service_sid"`
	MailGunServiceProtocol string `json:"mail_gun_service_protocol" xml:"mail_gun_service_protocol"`

	MailUser     string `json:"mail_user" xml:"mail_user"`
	MailPassword string `json:"mail_password" xml:"mail_password"`

	Env string `json:"env" xml:"env"`

	SaleNFTContractAddress721    string `json:"sale_nft_contract_address_721" xml:"sale_nft_contract_address_721"`
	AuctionNFTContractAddress721 string `json:"auction_nft_contract_address_721" xml:"auction_nft_contract_address_721"`
	MintNFTContractAddress721    string `json:"mint_nft_contract_address_721" xml:"mint_nft_contract_address_721"`
	BadgesContractAddress721     string `json:"badges_contract_address_721"xml:"badges_contract_address_721"`

	ElasticSearchURLs             []string `json:"elasticsearch_urls" xml:"elasticsearchurls"`
	ElasticSearchUserName         string   `json:"elasticsearch_username" xml:"elasticsearchusername"`
	ElasticSearchPassword         string   `json:"elasticsearch_password" xml:"elasticsearchpassword"`
	ElasticsearchIndexUsers       string   `json:"elasticsearch_index_users" xml:"elasticsearch_index_users"`
	ElasticsearchIndexItems       string   `json:"elasticsearch_index_items" xml:"elasticsearch_index_items"`
	ElasticsearchIndexCollections string   `json:"elasticsearch_index_collections" xml:"elasticsearch_index_collections"`
	ElasticsearchIndexEvents      string   `json:"elasticsearch_index_events" xml:"elasticsearch_index_events"`
	ElasticsearchIndexReports     string   `json:"elasticsearch_index_reports" xml:"elasticsearch_index_reports"`
	ElasticsearchIndexAuctions    string   `json:"elasticsearch_index_auctions" xml:"elasticsearch_index_auctions"`
	ElasticsearchIndexBids        string   `json:"elasticsearch_index_bids" xml:"elasticsearch_index_bids"`

	KafkaConsumerHost  string `json:"kafka_consumer_host" xml:"kafka_consumer_host"`
	KafkaConsumerPort  string `json:"kafka_consumer_port" xml:"kafka_consumer_port"`
	KafkaConsumerTopic string `json:"kafka_consumer_topic" xml:"kafka_consumer_topic"`

	KafkaProducerHost  string `json:"kafka_producer_host" xml:"kafka_producer_host"`
	KafkaProducerPort  string `json:"kafka_producer_port" xml:"kafka_producer_port"`
	KafkaProducerTopic string `json:"kafka_producer_topic" xml:"kafka_producer_topic"`

	KafkaBrokers  []string `json:"kafka_brokers" xml:"kafka_brokers"`
	KafkaProtocol string   `json:"kafka_protocol" xml:"kafka_protocol"`
}

var Config *AppConfig

func LoadInitConfig() {
	var err error
	config := AppConfig{}
	env := ""

	helps.ReadFlags(&env)
	log.Println(env, ": env")
	configPath := helps.ParseConfigPath(env)
	log.Println(configPath, "configPath appconfig/appconfig.go:31")
	err = beego.LoadAppConfig("ini", configPath)
	if err != nil {
		log.Println(err.Error(), "-- err.Error() appconfig/appconfig.go:34")
	}

	config.EtcdServerEndpoints, _ = beego.AppConfig.Strings("etcd::host")
	fmt.Println("[MAIN_LOAD]: Etcd servers:", config.EtcdServerEndpoints)
	config.BnbRPC, _ = beego.AppConfig.String("default::bscrpc")
	config.BnbRPCS, _ = beego.AppConfig.Strings("default::bnbrpcs")
	config.PoolCap, err = beego.AppConfig.Int("default::gethpool")
	if err != nil {
		log.Println(err.Error(), "err.Error() appconfig/appconfig.go:128")
		config.PoolCap = 100
	}
	config.WkPoolBsc, err = beego.AppConfig.Int("default::wkpoolBsc")
	if err != nil {
		log.Println(err.Error(), "err.Error() appconfig/appconfig.go:341")
		config.WkPoolBsc = 10
	}
	config.NumRetry, err = beego.AppConfig.Int("default::numRetry")
	if err != nil {
		log.Println(err.Error(), "err.Error() appconfig/appconfig.go:140")
		config.NumRetry = 2
	}

	config.TimeCache, _ = beego.AppConfig.Int64("default::timecache")

	config.Duration, _ = beego.AppConfig.Int64("default::duration")
	config.Env, _ = beego.AppConfig.String("default::env")

	config.IsSessionKey, _ = beego.AppConfig.Bool("default::issession")

	config.NFTThriftHost, _ = beego.AppConfig.String("nft_thrift_service::host")
	config.NFTThriftPort, _ = beego.AppConfig.String("nft_thrift_service::port")

	config.RedisHost, _ = beego.AppConfig.String("redis::host")
	config.RedisPassword, _ = beego.AppConfig.String("redis::pw")
	config.RedisPoolSize, _ = beego.AppConfig.Int("redis::poolsize")

	config.SessionServiceSid, _ = beego.AppConfig.String("sessionservice::sid")
	config.SessionServiceHost, _ = beego.AppConfig.String("sessionservice::host")
	config.SessionServicePort, _ = beego.AppConfig.String("sessionservice::port")
	config.SessionServiceProtocol, _ = beego.AppConfig.String("sessionservice::protocol")

	config.Pubkey2UidSid, _ = beego.AppConfig.String("pubkey2uid::sid")
	config.Pubkey2UidHost, _ = beego.AppConfig.String("pubkey2uid::host")
	config.Pubkey2UidPort, _ = beego.AppConfig.String("pubkey2uid::port")

	config.KvCounterSid, _ = beego.AppConfig.String("kvcounter::sid")
	config.KvCounterHost, _ = beego.AppConfig.String("kvcounter::host")
	config.KvCounterPort, _ = beego.AppConfig.String("kvcounter::port")

	config.NftMarketplaceSid, _ = beego.AppConfig.String("nft_marketplace::sid")
	config.NftMarketplaceHost, _ = beego.AppConfig.String("nft_marketplace::host")
	config.NftMarketplacePort, _ = beego.AppConfig.String("nft_marketplace::port")

	config.NftItemSid, _ = beego.AppConfig.String("nft_item::sid")
	config.NftItemHost, _ = beego.AppConfig.String("nft_item::host")
	config.NftItemPort, _ = beego.AppConfig.String("nft_item::port")

	config.NftItemMappingSid, _ = beego.AppConfig.String("nft_item_mapping::sid")
	config.NftItemMappingHost, _ = beego.AppConfig.String("nft_item_mapping::host")
	config.NftItemMappingPort, _ = beego.AppConfig.String("nft_item_mapping::port")

	config.NftUserSid, _ = beego.AppConfig.String("nft_user::sid")
	config.NftUserHost, _ = beego.AppConfig.String("nft_user::host")
	config.NftUserPort, _ = beego.AppConfig.String("nft_user::port")

	config.NftEventSid, _ = beego.AppConfig.String("nft_event::sid")
	config.NftEventHost, _ = beego.AppConfig.String("nft_event::host")
	config.NftEventPort, _ = beego.AppConfig.String("nft_event::port")

	config.Runmode, _ = beego.AppConfig.String("default::runmode")
	config.HttpHost, _ = beego.AppConfig.String("default::httphost")
	config.HttpPort, _ = beego.AppConfig.Int("default::httpport")

	config.MailUser, _ = beego.AppConfig.String("mail::user")
	config.MailPassword, _ = beego.AppConfig.String("mail::password")

	config.MailGunServiceHost, _ = beego.AppConfig.String("mail_gun_service::host")
	config.MailGunServicePort, _ = beego.AppConfig.String("mail_gun_service::port")
	config.MailGunServiceSid, _ = beego.AppConfig.String("mail_gun_service::sid")
	config.MailGunServiceProtocol, _ = beego.AppConfig.String("mail_gun_service::protocol")

	config.SaleNFTContractAddress721, _ = beego.AppConfig.String("nft::salecontractaddress721")
	config.MintNFTContractAddress721, _ = beego.AppConfig.String("nft::mintcontractaddress721")
	config.AuctionNFTContractAddress721, _ = beego.AppConfig.String("nft::auctioncontractaddress721")
	config.BadgesContractAddress721, _ = beego.AppConfig.String("nft::badgescontractaddress721")

	config.ElasticSearchURLs, _ = beego.AppConfig.Strings("elasticsearch::urls")
	config.ElasticSearchUserName, _ = beego.AppConfig.String("elasticsearch::username")
	config.ElasticSearchPassword, _ = beego.AppConfig.String("elasticsearch::password")
	config.ElasticsearchIndexUsers, _ = beego.AppConfig.String("elasticsearch::indexusers")
	config.ElasticsearchIndexItems, _ = beego.AppConfig.String("elasticsearch::indexitems")
	config.ElasticsearchIndexCollections, _ = beego.AppConfig.String("elasticsearch::indexcollections")
	config.ElasticsearchIndexEvents, _ = beego.AppConfig.String("elasticsearch::indexevents")
	config.ElasticsearchIndexReports, _ = beego.AppConfig.String("elasticsearch::indexreports")
	config.ElasticsearchIndexAuctions, _ = beego.AppConfig.String("elasticsearch::indexauctions")
	config.ElasticsearchIndexBids, _ = beego.AppConfig.String("elasticsearch::indexbids")

	config.KafkaBrokers, _ = beego.AppConfig.Strings("kafka::brokers")
	config.KafkaProtocol, _ = beego.AppConfig.String("kafka::protocol")

	config.KafkaConsumerHost, _ = beego.AppConfig.String("kafka_consumer::host")
	config.KafkaConsumerPort, _ = beego.AppConfig.String("kafka_consumer::port")
	config.KafkaConsumerTopic, _ = beego.AppConfig.String("kafka_consumer::topic")

	config.KafkaProducerHost, _ = beego.AppConfig.String("kafka_producer::host")
	config.KafkaProducerPort, _ = beego.AppConfig.String("kafka_producer::port")
	config.KafkaProducerTopic, _ = beego.AppConfig.String("kafka_producer::topic")

	log.Println("===== MintNFTContractAddress721 ===== : ", config.MintNFTContractAddress721)
	log.Println("===== SaleNFTContractAddress721 ===== : ", config.SaleNFTContractAddress721)
	log.Println("===== AuctionNFTContractAddress721 ===== : ", config.AuctionNFTContractAddress721)
	log.Println("===== BadgesContractAddress721 ===== : ", config.BadgesContractAddress721)

	Config = &config
}
