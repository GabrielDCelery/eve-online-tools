package custom

type MarketOrder struct {
	Duration     uint64  `json:"duration" csv:"duration"`
	IsBuyOrder   bool    `json:"is_buy_order" csv:"is_buy_order" parquet:"name=is_buy_order, type=BOOLEAN"`
	Issued       string  `json:"issued" csv:"issued"`
	LocationID   uint64  `json:"location_id" csv:"location_id"`
	MinVolume    uint64  `json:"min_volume" csv:"min_volume"`
	OrderID      uint64  `json:"order_id" csv:"order_id"`
	Price        float64 `json:"price" csv:"price" parquet:"name=price, type=FLOAT"`
	Range        string  `json:"range" csv:"range"`
	SystemID     uint64  `json:"system_id" csv:"system_id"`
	TypeID       uint64  `json:"type_id" csv:"type_id"`
	VolumeRemain uint64  `json:"volume_remain" csv:"volume_remain"`
	VolumeTotal  uint64  `json:"volume_total" csv:"volume_total"`
	DownloadedAt string  `json:"downloaded_at" csv:"downloaded_at"`
	RegionID     uint64  `json:"region_id" csv:"region_id"`
}

type MarketHistoryItem struct {
	Average      float64 `json:"average"`
	Date         string  `json:"date"`
	Highest      float64 `json:"highest"`
	Lowest       float64 `json:"lowest"`
	OrderCount   uint64  `json:"order_count"`
	Volume       uint64  `json:"volume"`
	DownloadedAt string  `json:"downloaded_at" csv:"downloaded_at"`
	TypeID       uint64  `json:"type_id" csv:"type_id"`
	RegionID     uint64  `json:"region_id" csv:"region_id"`
}
