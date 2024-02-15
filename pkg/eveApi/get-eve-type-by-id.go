package eveAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetEveTypeByID(typeID uint64) (EveApiType, error) {
	url := fmt.Sprintf(`https://esi.evetech.net/latest/universe/types/%d/?datasource=tranquility&language=en`, typeID)

	resp, api_error := http.Get(url)

	if api_error != nil {
		return EveApiType{}, api_error
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	eveEntity := EveApiType{}

	err = json.Unmarshal(body, &eveEntity)

	if err != nil {
		return EveApiType{}, err
	}

	return eveEntity, nil
}
