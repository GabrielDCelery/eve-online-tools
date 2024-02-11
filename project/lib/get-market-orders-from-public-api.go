package lib

import (
	"fmt"

	"github.com/GabrielDCelery/eve-online-tools-cli/pkg/dataTransforms"
	eveAPI "github.com/GabrielDCelery/eve-online-tools-cli/pkg/eveApi"

	"sync"
	"time"
)

func callAPI(config *GetMarketOrdersFromPublicApiConfig, mtx *sync.Mutex, toPrintChan *chan string, page *uint64) {
	eveApiMarketOrders, err := eveAPI.GetEveMarketOrders(config.RegionID, *page)
	downloadedAt := time.Now().Format("2006-01-02T15:04:05Z")

	if err != nil {
		close(*toPrintChan)
		fmt.Println(err)
		return
	}

	if len(eveApiMarketOrders) == 0 {
		close(*toPrintChan)
		return
	}

	go func() {
		for _, eveApiMarketOrder := range eveApiMarketOrders {
			internalMarketOrder := dataTransforms.TransformEveApiMarketOrderToMarketOrder(&eveApiMarketOrder, config.RegionID, downloadedAt)
			csvRow := dataTransforms.TransformMarketOrderToCsvRow(&internalMarketOrder)
			*toPrintChan <- csvRow
		}
	}()

	mtx.Lock()
	*page += 1
	mtx.Unlock()
	go callAPI(config, mtx, toPrintChan, page)

}

type GetMarketOrdersFromPublicApiConfig struct {
	RegionID uint64
}

func GetMarketOrdersFromPublicApi(config *GetMarketOrdersFromPublicApiConfig) {
	fmt.Println(dataTransforms.GetMarketOrderAsCsvHeader())

	mtx := sync.Mutex{}
	toPrintChan := make(chan string)

	var page uint64 = 1

	go callAPI(config, &mtx, &toPrintChan, &page)

	for toPrint := range toPrintChan {
		fmt.Println(toPrint)
	}
}
