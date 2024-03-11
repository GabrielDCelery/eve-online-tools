package lib

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/GabrielDCelery/eve-online-tools-cli/pkg/custom"
	"github.com/GabrielDCelery/eve-online-tools-cli/pkg/dataTransforms"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/writer"
)

type TransformMarketOrdersConfig struct {
	From string
	To   string
}

func TransforMarketOrders(config *TransformMarketOrdersConfig) {
	if config.From == "csv" {
		if config.To == "parquet" {

			reader := csv.NewReader(os.Stdin)

			fileWriter, err := local.NewLocalFileWriter("marketorders.parquet")

			if err != nil {
				log.Fatalln(err)
			}

			defer fileWriter.Close()

			parquetWriter, err := writer.NewParquetWriter(fileWriter, new(custom.MarketOrder), 1)

			if err != nil {
				log.Fatalln(err)
			}

			// parquetWriter.CompressionType = parquet.CompressionCodec_LZO
			// parquetWriter.RowGroupSize = 128 * 1024 * 1024 // 128MB

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

				marketOrder, err := dataTransforms.TransformCsvRowToMarketOrder(&csvRow)

				if err != nil {
					log.Fatalln(err)
				}

				if err = parquetWriter.Write(marketOrder); err != nil {
					log.Fatalln(err)
				}
			}

			if err := parquetWriter.WriteStop(); err != nil {
				log.Fatalln(err)
			}
		}
		return
	}

	log.Fatalln(fmt.Errorf("transformation from %s to %s not implemented yet", config.From, config.To))
}
