package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type marketInformationResponse struct {
	ProductCode string `json:"product_code"`
	MarketType  string `json:"market_type"`
	Alias       string `json:"alias,omitempty"`
}

func GetMarketInformation() (*[]marketInformationResponse, error) {
	res, err := http.Get("https://api.bitflyer.com/v1/markets")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var marketInformationResponses []marketInformationResponse

	if err := json.Unmarshal(body, &marketInformationResponses); err != nil {
		fmt.Println(err)
	}

	return &marketInformationResponses, nil
}
