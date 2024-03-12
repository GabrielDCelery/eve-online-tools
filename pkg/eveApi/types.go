package eveAPI

type EveApiMarketOrder struct {
	Duration     uint64  `json:"duration" csv:"duration"`
	IsBuyOrder   bool    `json:"is_buy_order" csv:"is_buy_order"`
	Issued       string  `json:"issued" csv:"issued"`
	LocationID   uint64  `json:"location_id" csv:"location_id"`
	MinVolume    uint64  `json:"min_volume" csv:"min_volume"`
	OrderID      uint64  `json:"order_id" csv:"order_id"`
	Price        float64 `json:"price" csv:"price"`
	Range        string  `json:"range" csv:"range"`
	SystemID     uint64  `json:"system_id" csv:"system_id"`
	TypeID       uint64  `json:"type_id" csv:"type_id"`
	VolumeRemain uint64  `json:"volume_remain" csv:"volume_remain"`
	VolumeTotal  uint64  `json:"volume_total" csv:"volume_total"`
}

type EveApDogmaAttribute struct {
	AttributeID uint64  `json:"attribute_id"`
	Volume      float64 `json:"value"`
}

type EveApiType struct {
	Capacity        uint64                `json:"capacity"`
	Description     string                `json:"description"`
	DogmaAttributes []EveApDogmaAttribute `json:"dogma_attributes"`
	GroupID         uint64                `json:"group_id"`
	IconID          uint64                `json:"icon_id"`
	MarketGroupID   uint64                `json:"market_group_id"`
	Mass            uint64                `json:"mass"`
	Name            string                `json:"name"`
	PackagedVolume  float64               `json:"packaged_volume"`
	PortionSize     uint64                `json:"portion_size"`
	Published       bool                  `json:"published"`
	Radius          uint64                `json:"radius"`
	TypeID          uint64                `json:"type_id"`
	Volume          float64               `json:"volume"`
}

type EveApiMarketHistoryItem struct {
	Average    float64 `json:"average"`
	Date       string  `json:"date"`
	Highest    float64 `json:"highest"`
	Lowest     float64 `json:"lowest"`
	OrderCount uint64  `json:"order_count"`
	Volume     uint64  `json:"volume"`
}
