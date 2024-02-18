package eveAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetEveMarketOrders(regionID uint64, page uint64) ([]EveApiMarketOrder, error) {
	url := fmt.Sprintf(`https://esi.evetech.net/latest/markets/%d/orders/?datasource=tranquility&order_type=all&page=%d`, regionID, page)

	resp, err := http.Get(url)

	if err != nil {
		return []EveApiMarketOrder{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			return []EveApiMarketOrder{}, err
		}

		eveApiMarketOrdersForPage := []EveApiMarketOrder{}

		err = json.Unmarshal(body, &eveApiMarketOrdersForPage)

		if err != nil {
			return []EveApiMarketOrder{}, err
		}

		return eveApiMarketOrdersForPage, nil
	}

	if resp.StatusCode == 404 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return []EveApiMarketOrder{}, err
		}
		if resp.StatusCode == 404 && string(body) == `{"error":"Requested page does not exist!"}` {
			return []EveApiMarketOrder{}, nil
		}
		return []EveApiMarketOrder{}, fmt.Errorf(`failed to call EVE API, status code %d`, resp.StatusCode)
	}

	return []EveApiMarketOrder{}, fmt.Errorf(`sorry something enexpected happened`)

}
