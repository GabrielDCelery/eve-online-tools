package commands

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"example/eve-online-tools/api"
	"log"
	"strconv"
)

type ExecuteGetMarketOrdersCommandParams struct {
	sectorID string
	format   string
}

func validateParams(paramsMap map[string]string) (ExecuteGetMarketOrdersCommandParams, error) {
	commandCompatibleParams := ExecuteGetMarketOrdersCommandParams{format: "json"}
	for k, v := range paramsMap {
		switch k {
		case "sectorID":
			commandCompatibleParams.sectorID = v
		case "format":
			switch v {
			case "json":
				commandCompatibleParams.format = v
			case "csv":
				commandCompatibleParams.format = v
			default:
				return ExecuteGetMarketOrdersCommandParams{}, &InvalidParameterValueError{Parameter: k, ParameterValue: v}
			}
		default:
			return ExecuteGetMarketOrdersCommandParams{}, &UnknownParameterError{Parameter: k}
		}
	}
	return commandCompatibleParams, nil
}

func ExecuteGetMarketOrdersCommand(paramsMap map[string]string) (string, error) {
	commandParams, err := validateParams(paramsMap)

	if err != nil {
		return "", err
	}

	eveMarketOrdersForSector, err := api.GetEveMarketOrdersForSector(commandParams.sectorID)

	if err != nil {
		return "", err
	}

	if commandParams.format == "csv" {
		var b bytes.Buffer
		csvWriter := csv.NewWriter(&b)

		csvHeader := []string{
			"duration",
			"is_buy_order",
			"issued",
			"location_id",
			"min_volume",
			"order_id",
			"price",
			"range",
			"system_id",
			"type_id",
			"volume_remain",
			"volume_total",
		}

		if err := csvWriter.Write(csvHeader); err != nil {
			log.Fatalln(err)
		}

		for _, marketOrder := range eveMarketOrdersForSector {
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
			)
			if err := csvWriter.Write(csvRow); err != nil {
				log.Fatalln(err)
			}
		}

		csvWriter.Flush()

		return b.String(), nil
	}

	final, err := json.Marshal(eveMarketOrdersForSector)

	if err != nil {
		log.Fatalln(err)
	}

	return string(final), nil
}
