package dataFilters

import "github.com/GabrielDCelery/eve-online-tools-cli/pkg/internals"

func HasSoldUnits(i *internals.MarketOrder) bool {
	return i.VolumeTotal > i.VolumeRemain
}

func DoesMatchTypeID(i *internals.MarketOrder, typeID uint64) bool {
	return i.TypeID == typeID
}
