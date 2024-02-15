package lib

import (
	"fmt"
	"log"
	"time"

	"github.com/GabrielDCelery/eve-online-tools-cli/pkg/dataTransforms"
	eveAPI "github.com/GabrielDCelery/eve-online-tools-cli/pkg/eveApi"
)

func GetMarketHistoryFromPublicApi(regionID uint64, typeID uint64) {
	downloadedAt := time.Now().Format("2006-01-02T15:04:05Z")

	eveMarketHistoryItems, err := eveAPI.GetEveMarketHistory(regionID, typeID)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(dataTransforms.GetMarketHistoryItemAsCsvHeader())

	for _, eveMarketHistoryItem := range eveMarketHistoryItems {
		marketHistoryItem := dataTransforms.TransformEveApiMarketHistoryItemToMarketHistoryItem(&eveMarketHistoryItem, regionID, typeID, downloadedAt)
		csvRow := dataTransforms.TransformMarketHistoryItemToCsvRow(&marketHistoryItem)
		fmt.Println(csvRow)
	}
}
