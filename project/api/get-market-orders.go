package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetEveMarketOrdersForSector(sectorID string) ([]EveMarketOrder, error) {
	page := 1
	keepRunning := true

	eveMarketOrdersForSector := []EveMarketOrder{}

	for keepRunning {
		url := fmt.Sprintf(`https://esi.evetech.net/latest/markets/%s/orders/?datasource=tranquility&order_type=all&page=%d`, sectorID, page)

		resp, err := http.Get(url)

		if err != nil {
			return []EveMarketOrder{}, err
		}

		if resp.StatusCode == 200 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return []EveMarketOrder{}, err
			}
			eveMarketOrdersForPage := []EveMarketOrder{}
			json.Unmarshal(body, &eveMarketOrdersForPage)
			eveMarketOrdersForSector = append(eveMarketOrdersForSector, eveMarketOrdersForPage...)
			page++
			continue
		}

		if resp.StatusCode == 404 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return []EveMarketOrder{}, err
			}
			if resp.StatusCode == 404 && string(body) == `{"error":"Requested page does not exist!"}` {
				keepRunning = false
				continue
			}
			return []EveMarketOrder{}, errors.New(fmt.Sprintf(`Failed to call EVE API, status code %d`, resp.StatusCode))
		}

		return []EveMarketOrder{}, errors.New(fmt.Sprintf(`Failed to call EVE API, status code %d`, resp.StatusCode))
	}

	return eveMarketOrdersForSector, nil

}
