package lib

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/GabrielDCelery/eve-online-tools-cli/pkg/dataCalculations"
	"github.com/GabrielDCelery/eve-online-tools-cli/pkg/dataTransforms"
)

func CalculateDailySale() {
	reader := csv.NewReader(os.Stdin)

	rowIndex := 0
	numOfRecords := 0
	totalNumOfDailySales := float64(0)

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

		numOfRecords += 1

		marketOrder, err := dataTransforms.TransformCsvRowToMarketOrder(&csvRow)

		if err != nil {
			log.Fatalln(err)
		}

		volumePerDay, err := dataCalculations.CalculateVolumePerDay(&marketOrder)

		if err != nil {
			log.Fatalln(err)
		}

		totalNumOfDailySales += volumePerDay
	}

	averageNumOfDailySales := totalNumOfDailySales / float64(numOfRecords)

	fmt.Println(averageNumOfDailySales)
}
