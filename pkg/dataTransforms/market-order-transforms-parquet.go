package dataTransforms

import "time"

const dateTimeFormatWithMillis = "2006-01-02T15:04:05Z"

func TransformMarketOrderToMarketOrderParquet(marketOrder *MarketOrder) (MarketOrderParquet, error) {
	issued, err := time.Parse(dateTimeFormatWithMillis, marketOrder.Issued)
	if err != nil {
		return MarketOrderParquet{}, err
	}
	downloadedAt, err := time.Parse(dateTimeFormatWithMillis, marketOrder.DownloadedAt)
	if err != nil {
		return MarketOrderParquet{}, err
	}
	marketOrderParquet := MarketOrderParquet{
		Duration:     int64(marketOrder.Duration),
		IsBuyOrder:   marketOrder.IsBuyOrder,
		Issued:       int32(issued.Unix()),
		LocationID:   int64(marketOrder.LocationID),
		MinVolume:    int64(marketOrder.MinVolume),
		OrderID:      int64(marketOrder.OrderID),
		Price:        marketOrder.Price,
		Range:        marketOrder.Range,
		SystemID:     int64(marketOrder.SystemID),
		TypeID:       int64(marketOrder.TypeID),
		VolumeRemain: int64(marketOrder.VolumeRemain),
		VolumeTotal:  int64(marketOrder.VolumeTotal),
		DownloadedAt: int32(downloadedAt.Unix()),
		RegionID:     int64(marketOrder.RegionID),
	}
	return marketOrderParquet, nil
}
