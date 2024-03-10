package dataTransforms

import (
	"strconv"
	"strings"

	"github.com/GabrielDCelery/eve-online-tools-cli/pkg/custom"
)

func TransformMarketHistoryItemToCsvRow(marketHistoryItem *custom.MarketHistoryItem) string {
	var csvRow []string
	csvRow = append(
		csvRow,
		marketHistoryItem.Date,
		strconv.FormatFloat(marketHistoryItem.Average, 'f', 0, 64),
		strconv.FormatFloat(marketHistoryItem.Highest, 'f', 0, 64),
		strconv.FormatFloat(marketHistoryItem.Lowest, 'f', 0, 64),
		strconv.FormatUint(marketHistoryItem.OrderCount, 10),
		strconv.FormatUint(marketHistoryItem.Volume, 10),
		marketHistoryItem.DownloadedAt,
		strconv.FormatUint(marketHistoryItem.TypeID, 10),
		strconv.FormatUint(marketHistoryItem.RegionID, 10),
	)
	return strings.Join(csvRow, ",")
}

func GetMarketHistoryItemAsCsvHeader() string {
	return `date,average,highest,lowest,order_count,volume,downloaded_at,type_id,region_id`
}
