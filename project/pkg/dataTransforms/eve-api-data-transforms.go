package dataTransforms

import (
	eveAPI "github.com/GabrielDCelery/eve-online-tools-cli/pkg/eveApi"
	"github.com/GabrielDCelery/eve-online-tools-cli/pkg/internals"
)

func TransformEveApiMarketOrderToMarketOrder(eMorder *eveAPI.EveApiMarketOrder, regionID uint64, downloadedAt string) internals.MarketOrder {
	return internals.MarketOrder{
		Duration:     eMorder.Duration,
		IsBuyOrder:   eMorder.IsBuyOrder,
		Issued:       eMorder.Issued,
		LocationID:   eMorder.LocationID,
		MinVolume:    eMorder.MinVolume,
		OrderID:      eMorder.OrderID,
		Price:        eMorder.Price,
		Range:        eMorder.Range,
		SystemID:     eMorder.SystemID,
		TypeID:       eMorder.TypeID,
		VolumeRemain: eMorder.VolumeRemain,
		VolumeTotal:  eMorder.VolumeTotal,
		DownloadedAt: downloadedAt,
		RegionID:     regionID,
	}
}

func TransformEveApiMarketHistoryItemToMarketHistoryItem(eMHItem *eveAPI.EveApiMarketHistoryItem, regionID uint64, typeID uint64, downloadedAt string) internals.MarketHistoryItem {
	return internals.MarketHistoryItem{
		Average:      eMHItem.Average,
		Date:         eMHItem.Date,
		Highest:      eMHItem.Highest,
		Lowest:       eMHItem.Lowest,
		OrderCount:   eMHItem.OrderCount,
		Volume:       eMHItem.Volume,
		DownloadedAt: downloadedAt,
		TypeID:       typeID,
		RegionID:     regionID,
	}
}
