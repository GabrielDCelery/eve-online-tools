package dataTransforms

type MarketOrderParquet struct {
	Duration     int64   `parquet:"name=duration, type=INT64, convertedtype=INT64"`
	IsBuyOrder   bool    `parquet:"name=is_buy_order, type=BOOLEAN"`
	Issued       int32   `parquet:"name=issued, type=INT32, convertedtype=TIME_MILLIS"`
	LocationID   int64   `parquet:"name=location_id, type=INT64, convertedtype=INT64"`
	MinVolume    int64   `parquet:"name=min_volume, type=INT64, convertedtype=INT64"`
	OrderID      int64   `parquet:"name=order_id, type=INT64, convertedtype=INT64"`
	Price        float64 `parquet:"name=price, type=DOUBLE"`
	Range        string  `parquet:"name=range, type=BYTE_ARRAY, convertedtype=UTF8"`
	SystemID     int64   `parquet:"name=system_id, type=INT64, convertedtype=INT64"`
	TypeID       int64   `parquet:"name=type_id, type=INT64, convertedtype=INT64"`
	VolumeRemain int64   `parquet:"name=volume_remain, type=INT64, convertedtype=INT64"`
	VolumeTotal  int64   `parquet:"name=volume_total, type=INT64, convertedtype=INT64"`
	DownloadedAt int32   `parquet:"name=downloaded_at, type=INT32, convertedtype=TIME_MILLIS"`
	RegionID     int64   `parquet:"name=region_id, type=INT64, convertedtype=INT64"`
}
