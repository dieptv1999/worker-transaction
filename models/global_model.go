package models

type ItemStatus int
type SaleType int
type EventType int
type CategoryType int
type MediaType int
type TypeUser int
type TypeReport int
type AuctionStatus int
type ReportStatus int

const (
	ItemCreated         ItemStatus = iota // bắt đầu tạo
	Freeze                                // đã đẩy có ipfs -> không cho sửa
	Pending                               // gửi lên cho admin kiểm duyệt
	AllowMint                             // admin cho phép đẩy mint blockchain
	NotAllowMint                          // admin không cho phép đẩy lên blockchain ( vi phạm )
	WaitBlockchain                        // đang đợi đẩy lên blockchain -> nếu ở trạng thái này không cho đẩy thêm nữa ( tránh việc đẩy trùng )
	Owned                                 // sau khi mint thành công
	Approved                              // item đã được approved
	OnSale                                // đang đăng bán
	OnAuction                             // đang đấu giá
	Deleted                               // đã bị xóa (burn khỏi blockchain)
	ItemInComingAuction                   // auction sắp diễn ra , cho cient filter - không update status này
	ItemAuctionEnded                      // auction đã kết thúc , cho cient filter - không update status này
)

const (
	AuctionCreated AuctionStatus = iota //  tạo phiên đấu giá
	AuctionEnded
)

const (
	Art CategoryType = iota
	Music
	DomainNames
	VirtualWorlds
	TradingCards
	Collectibles
	Sports
	Utility
	AllNFTs
	CategoryOther
)

const (
	Image MediaType = iota
	Video
	Audio
	Model3D
	MediaOther
)

const (
	NormalUser TypeUser = iota
	Artist
	Singer
	Gamer
)

func (s ItemStatus) String() string {
	return [...]string{"Created", "Freeze", "Pending", "AllowMint", "NotAllowMint", "WaitBlockchain", "Owned", "OnSale", "OnAuction", "Deleted"}[s]
}

func (p MediaType) String() string {
	switch p {
	case Image:
		return "Image"
	case Video:
		return "Video"
	case Audio:
		return "Audio"
	case Model3D:
		return "Model3D"
	case MediaOther:
		return "Other"
	default:
		return "Other"
	}
}

func (p TypeUser) String() string {
	switch p {
	case NormalUser:
		return "Normal User"
	case Artist:
		return "Artist"
	case Singer:
		return "Singer"
	case Gamer:
		return "Gamer"
	default:
		return "Normal User"
	}
}

func (p CategoryType) String() string {
	switch p {
	case Art:
		return "Art"
	case Music:
		return "Music"
	case DomainNames:
		return "Domain Names"
	case VirtualWorlds:
		return "Virtual Worlds"
	case TradingCards:
		return "Trading Cards"
	case Collectibles:
		return "Collectibles"
	case Sports:
		return "Sports"
	case Utility:
		return "Utility"
	case AllNFTs:
		return "AllNFTs"
	case CategoryOther:
		return "Other"
	default:
		return "Other"
	}
}

const (
	BuyNow SaleType = iota
	SellToTheHighestBidder
	SellWithDecliningPrice
)

func (p SaleType) String() string {
	switch p {
	case BuyNow:
		return "BuyNow"
	case SellToTheHighestBidder:
		return "SellToTheHighestBidder"
	case SellWithDecliningPrice:
		return "SellWithDecliningPrice"
	default:
		return "Unknown"
	}
}

const (
	EventTypeCreated EventType = iota
	EventTypeMinted
	EventTypeApproved
	EventTypeListingPrice
	EventTypeTransfer
	EventTypeBid
	EventTypeAuction
	EventTypeDeleted
	EventTypeCanceled
	EventTypeImported
	EventTypeClaimed
	EventTypeEditedAuction
	EventTypeCanceledAuction
)

func (p EventType) String() string {
	return [...]string{"EventTypeCreated", "EventTypeMint", "EventTypeApproved", "EventTypeListingPrice",
		"EventTypeTransfer", "EventTypeBid", "EventTypeAuction", "EventTypeDeleted", "EventTypeCanceled",
		"EventTypeImported", "EventTypeClaimed", "EventTypeEditedAuction", "EventTypeCanceledAuction"}[p]
}

const (
	ReportUser TypeReport = iota
	ReportItem
)

func (p TypeReport) String() string {
	return [...]string{"ReportUser", "ReportItem"}[p]
}

const (
	Waiting ReportStatus = iota
	Done
)

func (p ReportStatus) String() string {
	return [...]string{"Wait", "Done"}[p]
}
