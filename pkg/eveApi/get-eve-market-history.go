package eveAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetEveMarketHistory(regionID uint64, typeID uint64) ([]EveApiMarketHistoryItem, error) {
	url := fmt.Sprintf(`https://esi.evetech.net/latest/markets/%d/history/?datasource=tranquility&type_id=%d`, regionID, typeID)

	resp, err := http.Get(url)

	if err != nil {
		return []EveApiMarketHistoryItem{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return []EveApiMarketHistoryItem{}, err
	}

	var historyItems []EveApiMarketHistoryItem

	err = json.Unmarshal(body, &historyItems)

	if err != nil {
		return []EveApiMarketHistoryItem{}, err
	}

	return historyItems, nil
}
