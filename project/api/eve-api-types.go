package api

type EveMarketOrder struct {
	Duration     uint64  `json:"duration"`
	IsBuyOrder   bool    `json:"is_buy_order"`
	Issued       string  `json:"issued"`
	LocationID   uint64  `json:"location_id"`
	MinVolume    uint64  `json:"min_volume"`
	OrderID      uint64  `json:"order_id"`
	Price        float64 `json:"price"`
	Range        string  `json:"range"`
	SystemID     uint64  `json:"system_id"`
	TypeID       uint64  `json:"type_id"`
	VolumeRemain uint64  `json:"volume_remain"`
	VolumeTotal  uint64  `json:"volume_total"`
}
