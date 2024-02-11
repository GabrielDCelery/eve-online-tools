package lib

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/GabrielDCelery/eve-online-tools-cli/pkg/dataFilters"
	"github.com/GabrielDCelery/eve-online-tools-cli/pkg/dataTransforms"
)

func FilterMarketOrders(typeID uint64) {
	reader := csv.NewReader(os.Stdin)

	rowIndex := 0

	for {
		csvRow, err := reader.Read()

		rowIndex += 1

		if err != nil {
			switch err {
			case io.EOF:
				break
			default:
				log.Fatalln(err)
			}
		}

		if rowIndex == 1 {
			continue
		}

		if len(csvRow) == 0 {
			break
		}

		internalMarketOrder, err := dataTransforms.TransformCsvRowToInternalMarketOrder(&csvRow)

		if err != nil {
			log.Fatalln(err)
		}

		if dataFilters.HasSoldUnits(&internalMarketOrder) == false || dataFilters.DoesMatchTypeID(&internalMarketOrder, typeID) == false {
			continue
		}

		fmt.Println(internalMarketOrder)
	}
}
