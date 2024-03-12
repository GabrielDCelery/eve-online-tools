package dataTransforms

import (
	"strconv"
	"strings"
)

func TransformMarketOrderToCsvRow(marketOrder *MarketOrder) string {
	var csvRow []string
	csvRow = append(
		csvRow,
		strconv.FormatUint(marketOrder.Duration, 10),
		strconv.FormatBool(marketOrder.IsBuyOrder),
		marketOrder.Issued,
		strconv.FormatUint(marketOrder.LocationID, 10),
		strconv.FormatUint(marketOrder.MinVolume, 10),
		strconv.FormatUint(marketOrder.OrderID, 10),
		strconv.FormatFloat(marketOrder.Price, 'f', 0, 64),
		marketOrder.Range,
		strconv.FormatUint(marketOrder.SystemID, 10),
		strconv.FormatUint(marketOrder.TypeID, 10),
		strconv.FormatUint(marketOrder.VolumeRemain, 10),
		strconv.FormatUint(marketOrder.VolumeTotal, 10),
		marketOrder.DownloadedAt,
		strconv.FormatUint(marketOrder.RegionID, 10),
	)
	return strings.Join(csvRow, ",")
}

func GetMarketOrderAsCsvHeader() string {
	return `duration,is_buy_order,issued,location_id,min_volume,order_id,price,range,system_id,type_id,volume_remain,volume_total,downloaded_at,region_id`
}

func TransformCsvRowToMarketOrder(csvRow *[]string) (MarketOrder, error) {
	duration, err := strconv.ParseUint((*csvRow)[0], 10, 64)
	if err != nil {
		return MarketOrder{}, err
	}
	isByOrder, err := strconv.ParseBool((*csvRow)[1])
	if err != nil {
		return MarketOrder{}, err
	}
	issued := (*csvRow)[2]

	locationID, err := strconv.ParseUint((*csvRow)[3], 10, 64)
	if err != nil {
		return MarketOrder{}, err
	}
	minValue, err := strconv.ParseUint((*csvRow)[4], 10, 64)
	if err != nil {
		return MarketOrder{}, err
	}
	orderID, err := strconv.ParseUint((*csvRow)[5], 10, 64)
	if err != nil {
		return MarketOrder{}, err
	}
	price, err := strconv.ParseFloat((*csvRow)[6], 64)
	if err != nil {
		return MarketOrder{}, err
	}
	orderRange := (*csvRow)[7]

	systemID, err := strconv.ParseUint((*csvRow)[8], 10, 64)
	if err != nil {
		return MarketOrder{}, err
	}
	typeID, err := strconv.ParseUint((*csvRow)[9], 10, 64)
	if err != nil {
		return MarketOrder{}, err
	}
	volumeRemain, err := strconv.ParseUint((*csvRow)[10], 10, 64)
	if err != nil {
		return MarketOrder{}, err
	}
	volumeTotal, err := strconv.ParseUint((*csvRow)[11], 10, 64)
	if err != nil {
		return MarketOrder{}, err
	}
	downloadedAt := (*csvRow)[12]

	regionID, err := strconv.ParseUint((*csvRow)[13], 10, 64)
	if err != nil {
		return MarketOrder{}, err
	}

	marketOrder := MarketOrder{
		Duration:     duration,
		IsBuyOrder:   isByOrder,
		Issued:       issued,
		LocationID:   locationID,
		MinVolume:    minValue,
		OrderID:      orderID,
		Price:        price,
		Range:        orderRange,
		SystemID:     systemID,
		TypeID:       typeID,
		VolumeRemain: volumeRemain,
		VolumeTotal:  volumeTotal,
		DownloadedAt: downloadedAt,
		RegionID:     regionID,
	}

	return marketOrder, nil
}
